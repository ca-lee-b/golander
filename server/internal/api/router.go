package api

import "github.com/go-chi/chi/v5"

func (a *Api) initRoutes() {
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
