.PHONY: format build run test test-api test-service

format:
	go fmt ./...

build: format
	go build -o bin/time-api

run: build
	./bin/time-api

test:
	go test -count=1 ./tests/api ./tests/service

test-api:
	go test -count=1 ./tests/api

test-service:
	go test -count=1 ./tests/service
