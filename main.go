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

	myRouter := router.NewRouter("/")
	myRouter.Get("/", controller.IndexHandler)
	myRouter.Get("/sayHello", controller.HelloHandler)
	myRouter.Get("/api/json/hello", controller.HelloJsonResponse)
	myRouter.Get("/sayHello/:name", controller.HelloHandler)
	myRouter.Post("/auth", controller.AuthHandler)

	myRouter.Dispatch(context)
}
