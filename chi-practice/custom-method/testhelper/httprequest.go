package testhelper

import (
	"context"
	"io"
	"log"
	"net"
	"net/http"
	"time"
)

func HttpRequest(httpMethod string, url string, requestBody io.Reader) (int, []byte) {
	const ConnectMaxWaitTime = 1 * time.Second
	const RequestMaxWaitTime = 5 * time.Second
	cli := http.Client{
		Transport: &http.Transport{
			DialContext: (&net.Dialer{
				Timeout: ConnectMaxWaitTime,
			}).DialContext,
		},
	}
	ctx, cancel := context.WithTimeout(context.Background(), RequestMaxWaitTime)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, httpMethod, url, requestBody)
	if err != nil {
		log.Panicf("Cannot create request: %s\n", err)
	}

	log.Println("Do request")
	rsp, err := cli.Do(req)
	if rsp != nil {
		defer rsp.Body.Close()
	}
	if e, ok := err.(net.Error); ok && e.Timeout() {
		log.Panicf("Do request timeout: %s\n", err)
	} else if err != nil {
		log.Panicf("Cannot do request: %s\n", err)
	}

	log.Println("Read body")
	startRead := time.Now()
	body, err := io.ReadAll(rsp.Body)

	if err != nil {
		log.Panicf("Cannot read all response body: %s\n", err)
	}
	endRead := time.Now()
	log.Printf("Read response took: %s\n", endRead.Sub(startRead))

	log.Printf("%s\n", body)

	return rsp.StatusCode, body
}
