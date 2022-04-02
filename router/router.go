package router

import (
	"net/http"
)

// Router structure
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

func (r Router) Get(path string, handler http.HandlerFunc) {
	r.route(http.MethodGet, path, handler)
}

func (r Router) Post(path string, handler http.HandlerFunc) {
	r.route(http.MethodPost, path, handler)
}

func (r Router) Put(path string, handler http.HandlerFunc) {
	r.route(http.MethodPut, path, handler)
}

func (r Router) Delete(path string, handler http.HandlerFunc) {
	r.route(http.MethodDelete, path, handler)
}

func (r *Router) route(method string, path string, handler http.HandlerFunc) {
	r.routes = append(r.routes, Route{
		path, method, handler,
	})
}

// Route entry structure
type Route struct {
	Path    string
	Method  string
	Handler http.Handler
}

func (r *Route) Match(req *http.Request) bool {
	if req.Method != r.Method {
		return false
	}

	if req.URL.Path != r.Path {
		return false
	}

	return true
}
