package link

import (
	"custommethod/testhelper"
	"fmt"

	"net/http"
	"testing"
)

func TestLinkHandler(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name       string
		httpStatus int
	}{
		{name: "sample", httpStatus: http.StatusOK},
	}

	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			testServer := testhelper.RunTestServer("/", NewHandler)
			defer testServer.Close()

			status, _ := testhelper.HttpRequest(methodName, fmt.Sprintf("%s", testServer.URL), nil)
			if status != tt.httpStatus {
				t.Errorf("expected %v but got %d", tt.httpStatus, status)
			}
		})
	}
}
