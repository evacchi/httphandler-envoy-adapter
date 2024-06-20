package main

import (
	"net/http"

	"github.com/evacchi/httphandler-envoy-adapter/handler"
)

const Name = "headers"

func init() {
	handler.RegisterHandler(Name, http.HandlerFunc(myHttpHandler))
}

func myHttpHandler(w http.ResponseWriter, r *http.Request) {
	v := r.Header.Get("X-Custom-Header")
	if v == "greetings" {
		r.Header.Set("X-Hello", "hello to you!")
	}
}

func main() {}
