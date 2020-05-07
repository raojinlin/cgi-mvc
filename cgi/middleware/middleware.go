package middleware

import (
	"fmt"
	"iogo/cgi/controller"
	"iogo/cgi/http"
)

type Middleware func(next controller.Handler) controller.Handler
type Middlewares []Middleware

var Logger Middleware = func(next controller.Handler) controller.Handler {
	return func(ctx *http.Context) (int, string) {
		status, res := next(ctx)
		msg := fmt.Sprintf("%s %s %s %d %s %d\n",
			ctx.Request.RemoteAddr,
			ctx.Request.Method,
			ctx.Request.Uri,
			status,
			ctx.ServerProtocol,
			len(res),
		)

		if status >= 500 {
			ctx.Logger.Error(msg)
		} else {
			ctx.Logger.Info(msg)
		}

		return status, res
	}
}

