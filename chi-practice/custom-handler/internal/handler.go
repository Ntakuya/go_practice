package internal

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handler func(w http.ResponseWriter, r *http.Request) error

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		w.WriteHeader(503)
		w.Write([]byte("bad"))
	}
}

func customHandler(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query().Get("err")

	if q != "" {
		return errors.New(q)
	}

	w.Write([]byte("foo"))
	return nil
}

func InternalRouter(r chi.Router) chi.Router {
	r.Route("/api", func(r chi.Router) {
		r.Method("GET", "/", handler(customHandler))
	})
	return r
}
