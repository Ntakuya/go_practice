package woohoo

import (
	"custommethod/testhelper"
	"fmt"
	"net/http"
	"testing"
)

func TestWooHooHandler(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name         string
		httpStatus   int
		responseBody []byte
	}{
		{name: "sample", httpStatus: http.StatusOK, responseBody: []byte("custom woohoo method")},
	}
	for _, tt := range cases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			testServer := testhelper.RunTestServer("/", NewRoute)
			defer testServer.Close()

			status, body := testhelper.HttpRequest(NewMethod, fmt.Sprintf("%s", testServer.URL), nil)
			if status != tt.httpStatus {
				t.Errorf("status expected %v but got %v", tt.httpStatus, status)
			}
			if string(body) != string(tt.responseBody) {
				t.Errorf("body expected %v but got %v", string(tt.responseBody), string(body))
			}
		})
	}
}
