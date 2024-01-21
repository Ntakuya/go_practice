package helloworld

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func HelloWorldRoute(r chi.Router) {
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})
}
