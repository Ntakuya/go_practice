package customhandler

import (
	internal "customhandler/internal"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type app struct{}

func NewApp() *app {
	return &app{}
}

func (app *app) Run() {
	cnf := internal.NewConfig()
	r := chi.NewRouter()
	r.Method("GET", "/", internal.Handler(internal.CustomHandler))
	http.ListenAndServe(fmt.Sprintf(":%d", cnf.Port()), r)
}
