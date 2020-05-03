package main

import (
	"iogo/cgi"
	"iogo/cgi/config"
	"iogo/cgi/controller"
	"iogo/cgi/http"
	"iogo/cgi/router"
	"os"
)

func getConfigFile() string {
	configFile := "./config.json"
	if os.Getenv("IOGO_CONFIG_FILE") != "" {
		configFile = os.Getenv("IOGO_CONFIG_FILE")
	}

	return configFile
}

func getEnvs() map[string]string {
	if os.Getenv("IOGO_TEST_ENV") != "" {
		return cgi.GetEnvironMapFromEnvirons(cgi.Envs)
	}

	return cgi.GetEnvironMapFromEnvirons(os.Environ())
}

func main()  {
	cfg := config.Loader(getConfigFile())
	context := http.NewContext(getEnvs(), cfg)
	cgi.SetPostParamsToContext(context)

	myRouter := router.NewRouter("/")
	myRouter.Get("/", controller.IndexHandler)
	myRouter.Get("/sayHello", controller.HelloHandler)
	myRouter.Get("/api/json/hello", controller.HelloJsonResponse)
	myRouter.Get("/sayHello/:name", controller.HelloHandler)
	myRouter.Post("/auth", controller.AuthHandler)
	myRouter.Get("/403", controller.AccessDenied)
	myRouter.Get("/404", controller.NotFoundHandler)
	myRouter.Get("/500", controller.InternalServerErrorHandler)

	myRouter.Dispatch(context)
}
