# Http Handler Envoy Local Env

Example local development environment.

## Build

    docker compose -f docker/docker-compose-go.yaml run --rm go_plugin_compile

## Run

    docker compose -f docker/docker-compose.yaml up --build -d

Example:

    curl localhost:10000/localreply

Should reply:

    hello from path: /localreply

