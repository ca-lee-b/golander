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

	"github.com/go-chi/chi/v5"
)

type Api struct {
	Port string

	Router *chi.Mux
	Logger *slog.Logger
}

func New(port string, logger *slog.Logger) *Api {
	r := chi.NewRouter()

	return &Api{
		Port:   port,
		Router: r,
		Logger: logger,
	}
}

func (a *Api) Listen() error {
	formatted := fmt.Sprintf(":%v", a.Port)

	srv := &http.Server{
		Addr:    formatted,
		Handler: a.Router,
	}

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		err := srv.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic("Failed to start server")
		}
	}()

	<-shutdown
	a.Logger.Info("Shutdown signal recieved")
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := srv.Shutdown(ctx)
	if err != nil {
		return err
	}

	return nil
}
