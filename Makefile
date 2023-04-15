.PHONY: format build run serve test test-api test-service test-middleware coverage errcheck

format:
	go fmt ./...

build: format
	go build -o bin/time-api

run: build
	./bin/time-api

serve: run

test:
	go test -count=1 ./api/... ./middleware/...

test-api:
	go test -count=1 ./api/controllers

test-service:
	go test -count=1 ./api/services

test-middleware:
	go test -count=1 ./middleware/...

coverage:
	go test -cover ./...

errcheck:
	errcheck ./...
