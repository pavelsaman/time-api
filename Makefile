.PHONY: format build run serve test test-api test-service

format:
	go fmt ./...

build: format
	go build -o bin/time-api

run: build
	./bin/time-api

serve: run

test:
	go test -count=1 ./tests/api ./tests/service

test-api:
	go test -count=1 ./tests/api

test-service:
	go test -count=1 ./tests/service
