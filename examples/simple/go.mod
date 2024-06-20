module simple

go 1.20

replace github.com/evacchi/httphandler-envoy-adapter v0.0.0-20240619145623-613911cc8483 => ../..

replace github.com/envoyproxy/envoy => github.com/envoyproxy/envoy v1.30.1-0.20240618211207-87bc8a9bc8b3

require github.com/evacchi/httphandler-envoy-adapter v0.0.0-20240619145623-613911cc8483

require (
	github.com/envoyproxy/envoy v1.30.2 // indirect
	google.golang.org/protobuf v1.34.2 // indirect
)
