package internal

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
)

func runTestServer() *httptest.Server {
	r := chi.NewRouter()
	r.Group(func(r chi.Router) { InternalRouter(r) })
	return httptest.NewServer(r)
}

func TestInternalHandlerGet(t *testing.T) {
	cases := []struct {
		name   string
		status int
	}{
		{name: "should error", status: http.StatusNotFound},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			testServer := runTestServer()
			defer testServer.Close()
			resp, err := http.Get(fmt.Sprintf("%s/books", testServer.URL))
			// _, err := io.ReadAll(resp.Body)

			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != tt.status {
				t.Errorf("expected %d got: %v", tt.status, resp.StatusCode)
			}
		})
	}
}
