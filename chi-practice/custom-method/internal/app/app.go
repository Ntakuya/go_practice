package app

import (
	"custommethod/internal/app/helloworld"
	"custommethod/internal/app/link"
	"custommethod/internal/app/woohoo"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Run() {
	r := chi.NewRouter()
	setupServer(r)
}

func setupServer(r *chi.Mux) {
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Route("/", func(r chi.Router) {
		r.Route("/", helloworld.HelloWorldRoute)
		r.Route("/link", link.NewHandler)
		r.Route("/woo", woohoo.NewRoute)
		r.HandleFunc("/everything", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("capturing all standard http methods, as well as LINK, UNLINK and WOOHOO"))
		})
	})
	http.ListenAndServe(":3333", r)
}
