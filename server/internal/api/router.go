package api

func (a *Api) initRoutes() {
	a.Router.Get("/", a.UserHandler.GetOneById)
}
