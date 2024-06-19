package main

import (
	"fmt"
	"github.com/envoyproxy/envoy/contrib/golang/common/go/api"
	"net/http"
)

var UpdateUpstreamBody = "upstream response body updated by the simple plugin"

// The callbacks in the filter, like `DecodeHeaders`, can be implemented on demand.
// Because api.PassThroughStreamFilter provides a default implementation.
type filter struct {
	api.PassThroughStreamFilter

	callbacks api.FilterCallbackHandler
	config    *config

	handler        http.Handler
	request        *http.Request
	responseWriter *responseWriter
}

// Callbacks which are called in request path
// The endStream is true if the request doesn't have body
func (f *filter) DecodeHeaders(header api.RequestHeaderMap, endStream bool) api.StatusType {

	r := headersToRequest(header)
	rw := newResponseWriter()
	f.request, f.responseWriter = r, rw

	f.handler.ServeHTTP(rw, r)
	grpcStatus := int64(0)
	// If the ResponseWriter has been modified, then we can handle the response locally.
	if rw.modified {
		f.callbacks.DecoderFilterCallbacks().SendLocalReply(rw.status, rw.buf.String(), rw.header, grpcStatus, "")
		return api.LocalReply
	}
	// Write back changes in the request into the headers.
	requestToHeaders(r, header)
	return api.Continue
}

func headersToRequest(header api.RequestHeaderMap) *http.Request {
	url := extractUrl(header)

	r, err := http.NewRequest(header.Method(), url, nopReader{})
	if err != nil {
		panic(err)
	}

	r.Header = header.GetAllHeaders()
	return r
}

func extractUrl(header api.RequestHeaderMap) string {
	return fmt.Sprintf("%s://%s%s", header.Scheme(), header.Host(), header.Path())
}

func requestToHeaders(req *http.Request, header api.RequestHeaderMap) {
	header.SetHost(req.URL.Host)
	header.SetPath(req.URL.Path)
	header.SetMethod(req.Method)

	for k, v := range req.Header {
		header.Del(k)
		for i, vv := range v {
			if i == 0 {
				header.Set(k, vv)
			} else {
				header.Add(k, vv)
			}
		}
	}

}
