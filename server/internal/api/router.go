package api

func (a *Api) initRoutes() {
	a.Router.Get("/{id}", a.UserHandler.GetOneById)

	a.Router.Get("/events/", a.EventHandler.GetOneById)
}
