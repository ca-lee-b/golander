package api

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ca-lee-b/golander/server/internal/repositories"
	"github.com/go-chi/chi/v5"
)

type Api struct {
	Port        string
	UserHandler *UserHandler

	Router *chi.Mux
	Logger *slog.Logger
}

func New(port string, repos *repositories.Repositories, logger *slog.Logger) *Api {
	r := chi.NewRouter()

	return &Api{
		Port:        port,
		UserHandler: newUserHandler(repos.UserRepository),
		Router:      r,
		Logger:      logger,
	}
}

func (a *Api) Listen() error {
	a.initRoutes()
	formatted := fmt.Sprintf(":%v", a.Port)

	srv := &http.Server{
		Addr:    formatted,
		Handler: a.Router,
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		a.Logger.Info(fmt.Sprintf("Listening on port %v", a.Port))
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic("Failed to start server")
		}
	}()

	<-shutdown
	a.Logger.Info("Beginning graceful shutdown")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}
