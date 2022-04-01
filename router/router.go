package router

import (
	"net/http"
)

type Router struct {
	routes []Route
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for _, route := range r.routes {
		if !route.Match(req) {
			continue
		}
		route.Handler.ServeHTTP(w, req)
		return
	}

	// Return 404
	http.NotFound(w, req)
}

func (r *Router) Route(method string, path string, handler http.HandlerFunc) {
	route := Route{
		path,
		method,
		handler,
	}
	r.routes = append(r.routes, route)
}
