# Http Handler Envoy Local Env

This is a PoC for an adapter from the [Go HTTP Filter for Envoy][go-http-filter]
to [`net/http` middleware (`Handler`)][go-mw] interface.

This is a self-contained developer environment, replicating and based on
[the Envoy sandbox][go-http-filter].

[go-http-filter]: https://www.envoyproxy.io/docs/envoy/latest/start/sandboxes/golang-http
[go-mw]: https://pkg.go.dev/net/http#Handler

## Build

    make build

Defaults to building `example/simple`. If you want to build a different plugin, e.g. `headers`, instead:

    make build EXAMPLE=headers

Currently, you will also have to update the `docker/envoy.yaml` section, setting the `plugin_name` accordingly. 

## Run

    make up

This builds the Envoy container and an echo service (found under `docker/testing/echo.go`)
which will print out the contents of the incoming request in JSON format. 
Envoy listens on port 10000. For instance:

    $ curl localhost:10000 | jq

    {
      "Datetime": "...",
      "Method": "GET",
      "Path": "/",
      "Query": {},
      "Headers": {
        "Accept": [
          "*/*"
        ],
        ...
      },
      "Body": ""
    }

You can tear down all the created containers with

    make down


### Example: Simple

This plugin intercepts the `/localreply` path, and replies at the proxy level.

    curl localhost:10000/localreply

Should reply:

    hello from path: /localreply


### Example: Headers

This plugin inspects all requests for the header `X-Custom-Header`.
If `X-Custom-Header` contains the value `greetings`, then it will
inject an additional header `X-Hello: hello to you!`.

    curl localhost:10000 -H 'X-Custom-Header: greetings' | jq

Should reply:

```json
{
  ...
  "Headers": {
    "Accept": [
      "*/*"
    ],
    "X-Custom-Header": [
      "greetings"
    ],
    ...
    "X-Hello": [
      "hello to you!"
    ],
    ...
  },
}
```

## Writing your own

Create a `main` module, define an `http.Handler` as usual, then register it using the `handler.RegisterHandler()` API
in the `init()`. For instance:

```go
package main

import (
	"net/http"
	"github.com/evacchi/httphandler-envoy-adapter/handler"
)

func init() { handler.RegisterHandler("hello", http.HandlerFunc(hello)) }

// hello unconditionally sets the `X-Hello` header on all incoming requests.
func hello(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("X-Hello", "hello")
}

func main() {}
```
