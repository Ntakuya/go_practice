package link

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

const methodName = "LINK"

func NewHandler(r chi.Router) {
	chi.RegisterMethod(methodName)
	r.MethodFunc(methodName, "/", linkHandler)
}

func linkHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("custom link method"))
}
