package main

import (
	"net/http"

	"github.com/thrashzone13/go-router/router"
)

func main() {
	r := &router.Router{}
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello!"))
	})
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Pong!"))
	})
	http.ListenAndServe(":8000", r)
}
