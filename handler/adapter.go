package handler

import (
	"bytes"
	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
	ehttp "github.com/envoyproxy/envoy/contrib/golang/filters/http/source/go/pkg/http"
	"io"
	"net/http"
)

func RegisterHandler(name string, handler http.Handler) {
	ehttp.RegisterHttpFilterFactoryAndConfigParser(name, filterFactory(handler), ehttp.NullParser)
}

func filterFactory(handler http.Handler) api.StreamFilterFactory {
	return func(c interface{}, callbacks api.FilterCallbackHandler) api.StreamFilter {
		return &filter{
			callbacks: callbacks,
			handler:   handler,
		}
	}
}

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
