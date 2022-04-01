package router

import "net/http"

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
