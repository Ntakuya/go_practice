package helloworld

import (
	"custommethod/testhelper"
	"net/http"
	"testing"
)

func TestGetHandler(t *testing.T) {
	cases := []struct {
		name   string
		status int
	}{
		{name: "success", status: http.StatusOK},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			testServer := testhelper.RunTestServer("/", HelloWorldRoute)
			defer testServer.Close()
		})
	}
}
