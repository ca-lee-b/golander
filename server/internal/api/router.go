package api

func (a *Api) initRoutes() {
	a.Router.Get("/users/{id}", a.UserHandler.GetOneById)

	a.Router.Get("/events/{id}", a.EventHandler.GetOneById)
	a.Router.Post("/events", a.EventHandler.CreateEvent)
	a.Router.Post("/events/{id}", a.EventHandler.AddParticipant)
}
