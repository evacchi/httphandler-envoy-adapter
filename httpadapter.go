package main

import (
	gohttpapi "github.com/envoyproxy/envoy/contrib/golang/common/go/api"
	"net/http"
)

type envoyAdapter struct {
	gohttpapi.PassThroughStreamFilter
	handler http.Handler
}

func (e *envoyAdapter) DecodeHeaders(gohttpapi.RequestHeaderMap, bool) gohttpapi.StatusType {
	return gohttpapi.Continue
}

func (*envoyAdapter) DecodeData(gohttpapi.BufferInstance, bool) gohttpapi.StatusType {
	return gohttpapi.Continue
}

func (*envoyAdapter) DecodeTrailers(trailerMap gohttpapi.RequestTrailerMap) gohttpapi.StatusType {
	return gohttpapi.Continue
}

func main() {}
