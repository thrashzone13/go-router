package main

import (
	"net/http"

	"github.com/thrashzone13/go-router/router"
)

func main() {
	r := &router.Router{}
	r.Route(http.MethodGet, "/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	})
	r.Route(http.MethodGet, "/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong!"))
	})
	http.ListenAndServe(":8000", r)
}
