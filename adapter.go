package main

import (
	"bytes"
	"io"
	"net/http"
)

type nopReader struct{}

func (nopReader) Read(p []byte) (n int, err error) {
	return 0, io.EOF
}

type responseWriter struct {
	header   http.Header
	buf      bytes.Buffer
	status   int
	modified bool
}

func newResponseWriter() *responseWriter {
	return &responseWriter{header: http.Header{}, status: http.StatusOK}
}

func (r *responseWriter) Header() http.Header {
	r.modified = true
	return r.header
}

func (r *responseWriter) Write(bytes []byte) (int, error) {
	r.modified = true
	return r.buf.Write(bytes)
}

func (r *responseWriter) WriteHeader(statusCode int) {
	r.modified = true
	r.status = statusCode
}
