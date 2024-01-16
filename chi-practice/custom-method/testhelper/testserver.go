package testhelper

import (
	"net/http/httptest"

	"github.com/go-chi/chi/v5"
)

type chiRoute func(r chi.Router)

func RunTestServer(path string, fn chiRoute) *httptest.Server {
	r := chi.NewRouter()
	r.Route(path, fn)
	return httptest.NewServer(r)
}
