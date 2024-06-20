all:
	echo 'Available targets: build up'

.PHONY: build up down

build:
	docker compose -f docker/docker-compose-go.yaml run --rm go_plugin_compile

up:
	docker compose -f docker/docker-compose.yaml up --build

down:
	docker compose -f docker/docker-compose.yaml down
