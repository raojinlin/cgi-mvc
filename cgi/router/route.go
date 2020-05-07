package router

import (
	"iogo/cgi/controller"
	"iogo/cgi/http"
	"iogo/cgi/middleware"
)

type Route struct {
	route string
	handler controller.Handler
}

type Router struct {
	root string
	routes []*Route
	handlers middleware.Middlewares
}

func (r *Router) Use(handler ...middleware.Middleware) *Router {
	r.handlers = append(r.handlers, handler...)
	return r
}

func (r *Router) applyMiddleware(handler controller.Handler) controller.Handler {
	for i := 0; i < len(r.handlers); i++ {
		handler = r.handlers[i](handler)
	}

	return handler
}

func (r *Router) Get(path string, handler controller.Handler)  {
	r.routes = append(r.routes, &Route{
		route:   getRouteKey("get", path, r.root),
		handler: r.applyMiddleware(handler),
	})
}

func (r *Router) Post(path string, handler controller.Handler)  {
	r.routes = append(r.routes, &Route{
		route:   getRouteKey("post", path, r.root),
		handler: r.applyMiddleware(handler),
	})
}

func (r *Router) Dispatch(ctx *http.Context) {
	var handler controller.Handler
	var key = ctx.Request.Method + "#" + ctx.Request.Uri

	handler = r.applyMiddleware(controller.NotFoundHandler)
	for _, route := range r.routes {
		if route != nil && routeCompare(route.route, key) {
			handler = route.handler
			ctx.Request.SetParams(getRouteParams(route.route, key))
		}
	}

	c := controller.NewController(ctx, handler)
	c.Output()
}

func NewRouter(root string) *Router {
	return &Router{
		root:   root,
		routes: make([]*Route, 0),
		handlers: make(middleware.Middlewares, 0),
	}
}
