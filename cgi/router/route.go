package router

import (
	"iogo/cgi/controller"
	"iogo/cgi/http"
)

type Route struct {
	route string
	handler controller.Handler
}

type Router struct {
	root string
	routes []*Route
}

func (r *Router) Get(path string, handler controller.Handler)  {
	r.routes = append(r.routes, &Route{
		route:   getRouteKey("get", path, r.root),
		handler: handler,
	})
}

func (r *Router) Post(path string, handler controller.Handler)  {
	r.routes = append(r.routes, &Route{
		route:   getRouteKey("post", path, r.root),
		handler: handler,
	})
}

func (r *Router) Dispatch(ctx *http.Context) {
	var handler controller.Handler
	var key = ctx.Request.Method + "#" + ctx.Request.Uri

	handler = controller.NotFoundHandler
	for _, route := range r.routes {
		if route != nil && routeCompare(route.route, key) {
			handler = route.handler
			ctx.Request.SetParams(getRouteParams(route.route, key))
		}
	}

	controller.NewController(ctx, handler).Output()
}

func NewRouter(root string) *Router {
	return &Router{
		root:   root,
		routes: make([]*Route, 0),
	}
}
