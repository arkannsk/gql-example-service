.PHONY: setup build run gen-graphql

setup:
	cp env.sample .env
	cp docker/docker-compose.yml.dev  docker/docker-compose.yml

	mkdir -p jwt-keys
	openssl genrsa -out jwt-keys/jwt-private.key 4096
	openssl rsa -in jwt-keys/jwt-private.key -pubout -out jwt-keys/jwt-public.key

build: gen-graphql
	go build -mod vendor -o bin/server ./cmd/server/main.go

run:
	go run -mod vendor ./cmd/server/main.go

gen-graphql:
	gqlgen

