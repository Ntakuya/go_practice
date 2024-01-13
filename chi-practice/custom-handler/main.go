package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler func(w http.ResponseWriter, r *http.Request) error

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := h(w, r); err != nil {
		w.WriteHeader(503)
		w.Write([]byte("bad"))
	}
}

func main() {
	cnf := NewConfig()
	r := chi.NewRouter()
	r.Method("GET", "/", Handler(customHandler))
	http.ListenAndServe(fmt.Sprintf(":%d", cnf.Port()), r)
}

func customHandler(w http.ResponseWriter, r *http.Request) error {
	q := r.URL.Query().Get("err")

	if q != "" {
		return errors.New(q)
	}

	w.Write([]byte("foo"))
	return nil
}

type Config struct {
	port int
}

func NewConfig() *Config {
	return &Config{
		port: 3333,
	}
}

func (c *Config) Port() int {
	return c.port
}
