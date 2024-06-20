package main

import (
	"fmt"
	"net/http"

	"github.com/evacchi/httphandler-envoy-adapter/handler"
)

const Name = "simple"

func init() {
	handler.RegisterHandler(Name, http.HandlerFunc(myHttpHandler))
}

func myHttpHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/localreply" {
		body := fmt.Sprintf("hello from path: %s\r\n", path)
		w.Write([]byte(body))
	}
}

func main() {}
