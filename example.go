package main

import (
	"fmt"
	"net/http"
)

func myHttpHandler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	if path == "/localreply" {
		body := fmt.Sprintf("hello from path: %s\r\n", path)
		w.Write([]byte(body))
	}
}
