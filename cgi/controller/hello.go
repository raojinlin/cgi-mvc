package controller

import (
	"iogo/cgi/http"
	"time"
)

func HelloHandler(ctx *http.Context) (int, string) {
	if ctx.Request.Params["name"] != "" {
		return 200, "<h1>hello, " + ctx.Request.Params["name"] + "</h1>"
	}
	return 200, "<h1>hello world</h1>"
}

func HelloJsonResponse(ctx *http.Context) (int, string)  {
	return JsonOutput(ctx, "{\"message\": \"hello world\", \"status\": \"ok\"}")
}

func JsonOutput(ctx *http.Context, body string) (int, string) {
	ctx.Response.SetContentType("application/json")
	return 200, body
}

func WhatTimeIsNow(ctx *http.Context) (int, string) {
	return JsonOutput(ctx, "{\"time\": \"" + time.Now().String() + "\"}")
}