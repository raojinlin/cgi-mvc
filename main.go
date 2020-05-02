package main

import (
	"iogo/cgi"
	"iogo/cgi/config"
	"iogo/cgi/controller"
	"iogo/cgi/http"
	"iogo/cgi/router"
	"os"
)


func main()  {
	cfg := config.NewConfig("/home/raojinlin/GolandProjects/src/iogo/cgi/templates/")
	context := http.NewContext(cgi.GetEnvironMapFromEnvirons(os.Environ()), cfg)
	cgi.SetPostParamsToContext(context)

	myRouter := router.NewRouter("/cgi-bin/iogo")
	myRouter.Get("/", controller.IndexHandler)
	myRouter.Post("/auth", controller.AuthHandler)

	myRouter.Dispatch(context)
}
