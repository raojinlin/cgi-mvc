package main

import (
	"iogo/cgi"
	"iogo/cgi/controller"
)

func main()  {
	app := cgi.NewApplication("/")

	app.RouteGet("/", controller.IndexHandler)
	app.RouteGet("/sayHello", controller.HelloHandler)
	app.RouteGet("/api/json/hello", controller.HelloJsonResponse)
	app.RouteGet("/sayHello/:name", controller.HelloHandler)
	app.RoutePost("/auth", controller.AuthHandler)
	app.RouteGet("/403", controller.AccessDenied)
	app.RouteGet("/404", controller.NotFoundHandler)
	app.RouteGet("/500", controller.InternalServerErrorHandler)
	app.RouteGet("/api/time/now", controller.WhatTimeIsNow)

	app.Run()
}
