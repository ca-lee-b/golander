package api

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
)

func (a *Api) initRoutes() {
	a.Router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{os.Getenv("FRONTEND_URL"), "http://localhost:8080"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowCredentials: true,
	}))

	a.Router.Post("/login", a.AuthHandler.Login)
	a.Router.Post("/register", a.AuthHandler.Register)

	a.Router.Group(func(r chi.Router) {
		r.Use(a.AuthHandler.WithToken)
		r.Get("/users/{id}", a.UserHandler.GetOneById)
		r.Get("/events/{id}", a.EventHandler.GetOneById)
		r.Post("/events", a.EventHandler.CreateEvent)
		r.Post("/events/{id}", a.EventHandler.AddParticipant)
	})
}
