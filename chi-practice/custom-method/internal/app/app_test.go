package app

import (
	"custommethod/internal/app/helloworld"
	"custommethod/testhelper"
	"fmt"
	"net/http"
	"testing"
)

func TestCustomMethod(t *testing.T) {
	cases := []struct {
		name   string
		status int
	}{
		{name: "sample", status: http.StatusOK},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			testServer := testhelper.RunTestServer("/", helloworld.HelloWorldRoute)
			defer testServer.Close()
			resp, err := http.Get(fmt.Sprintf("%s", testServer.URL))
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
