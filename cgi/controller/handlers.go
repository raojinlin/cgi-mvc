package controller

import (
	"iogo/cgi/http"
	view2 "iogo/cgi/view"
	"time"
)

func NotFoundHandler(ctx *http.Context) (int, string) {
	return 404, view2.NewErrorView(ctx).Render("404", "Page Not Found.", 404)
}

func InternalServerErrorHandler(ctx *http.Context) (int, string) {
	return 500, view2.NewErrorView(ctx).Render("Internal Server Error.", "Internal Server Error.", 500)
}

func AccessDenied(ctx *http.Context) (int, string) {
	view := view2.NewErrorView(ctx)
	return 403, view.Render("Forbidden", "Access Denied.", 403)
}


func IndexHandler(ctx *http.Context) (int, string) {
	view := view2.NewIndexView(ctx)

	data := view2.IndexData{Name: "hello"}
	return 200, view.Render(data)
}

func AuthHandler(ctx *http.Context) (int, string) {
	sessionCookie := ctx.Cookie("session")
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
		cookie := http.NewCookie(ctx.Request.RemoteAddr)
		cookie.HttpOnly = true
		cookie.Path = "/"
		cookie.Expires = &tm
		cookie.Key = "session"
		cookie.Value = "yes"

		ctx.Response.SetCookie(cookie)

		ctx.Response.Redirect("/index.html")
		return 301, "ok"
	}

	return AccessDenied(ctx)
}