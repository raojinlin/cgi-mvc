package controller

import "iogo/cgi/http"

func HelloHandler(ctx *http.Context) (int, string) {
	if ctx.Request.Params["name"] != "" {
		return 200, "<h1>hello, " + ctx.Request.Params["name"] + "</h1>"
	}
	return 200, "<h1>hello world</h1>"
}

func HelloJsonResponse(ctx *http.Context) (int, string)  {
	ctx.Response.SetHeader("Content-Type", "application/json")
	return 200, "{\"message\": \"hello world\", \"status\": \"ok\"}"
}
