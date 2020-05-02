package controller

import (
	"iogo/cgi/http"
	view2 "iogo/cgi/view"
	"time"
)

func NotFoundHandler(ctx *http.Context) (int, string) {
	return 404, "<h1>404</h1> \n<p>Page Not Found.</p>\n " + ctx.ServerSignature
}

func InternalServerErrorHandler(ctx *http.Context) (int, string) {
	return 500, "<h1>500</h1> \n<p>Internal Server Error.</p>\n " + ctx.ServerSignature
}

func AccessDenied(ctx *http.Context) (int, string) {
	return 403, "<h1>403</h1>\n<p>Access Denied.</p>\n" + ctx.ServerSignature
}


func IndexHandler(ctx *http.Context) (int, string) {
	view := view2.NewIndexView(ctx)

	data := view2.IndexData{Name: "hello"}
	return 200, view.Render(data)
}

func AuthHandler(ctx *http.Context) (int, string) {
	sessionCookie := ctx.Request.GetCookie("session")
	if sessionCookie != nil && sessionCookie.Value == "yes" {
		ctx.Response.Redirect("/index.html?name=123")
		ctx.Response.SetCookie(&http.Cookie{
			Key:   "session",
			Value: "deleted",
		})
		return 301, "Location: /"
	}

	if ctx.Request.PostPrams["username"] == "user" && ctx.Request.PostPrams["password"] == "pass" {
		tm := time.Now()
		ctx.Response.SetCookie(&http.Cookie{
			Key:      "session",
			Value:    "yes",
			Domain:   "127.0.0.1",
			Path:     "/",
			Expires:  &tm,
			MaxAge:   0,
			HttpOnly: false,
			Secure:   false,
		})

		ctx.Response.Redirect("/index.html")
		return 301, "ok"
	}

	return AccessDenied(ctx)
}