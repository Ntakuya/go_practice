package woohoo

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

const NewMethod = "WOOHOO"

func NewRoute(r chi.Router) {
	chi.RegisterMethod(NewMethod)
	r.MethodFunc(NewMethod, "/", woohoohandler)
}

func woohoohandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("custom woohoo method"))
}
