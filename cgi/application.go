package cgi

import (
	"iogo/cgi/config"
	"iogo/cgi/controller"
	"iogo/cgi/http"
	"iogo/cgi/middleware"
	"iogo/cgi/router"
)

type Application struct {
	config *config.Config
	router *router.Router
	root string
}

func (app *Application) init()  {
	app.config = config.Loader(GetConfigFile())
	app.router = router.NewRouter(app.root)
}

func (app *Application) Run()  {
	ctx := http.NewContext(GetEnvs(), app.config)
	SetPostParamsToContext(ctx)
	app.router.Dispatch(ctx)
}

func (app *Application) RouteGet(path string, handler controller.Handler)  {
	app.router.Get(path, handler)
}

func (app *Application) RoutePost(path string, handler controller.Handler)  {
	app.router.Post(path, handler)
}

func (app *Application) Route() *router.Router  {
	return app.router
}

func NewApplication(root string) *Application  {
	app := &Application{root: root}
	app.init()
	app.router.Use(middleware.Logger)

	return app
}