package main

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Request struct {
	Datetime string
	Method   string
	Path     string
	Query    map[string][]string
	Headers  map[string][]string
	Body     string
}

func echo(w http.ResponseWriter, req *http.Request) {

	r := Request{
		Datetime: time.Now().UTC().String(),
		Method:   req.Method,
		Path:     strings.Split(req.URL.String(), "?")[0],
		Query:    req.URL.Query(),
		Headers:  req.Header,
		Body:     "",
	}
	if req.Body != nil {
		bodyBytes, err := io.ReadAll(req.Body)
		if err != nil {
			log.Printf("Body reading error: %v", err)
			return
		}
		r.Body = string(bodyBytes)
		defer req.Body.Close()
	}
	rb, _ := json.Marshal(r)

	_, ok := r.Query["delay"]
	if ok && len(r.Query["delay"]) > 0 {
		delay, _ := time.ParseDuration(r.Query["delay"][0] + "s")
		time.Sleep(delay)
	}

	w.Write(rb)
}

func setupEcho(ctx context.Context, port int) {
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if err != nil {
		log.Fatal(err)
	}
	srv := &http.Server{
		Handler: http.HandlerFunc(echo),
	}
	go func() {
		if err := srv.Serve(lis); err != nil {
			log.Fatal(err)
		}
	}()
	<-ctx.Done()
	srv.Close()
}
