.PHONY: format build run serve test test-api test-service

format:
	go fmt ./...

build: format
	go build -o bin/time-api

run: build
	./bin/time-api

serve: run

test:
	go test -count=1 ./api/controllers ./api/services

test-api:
	go test -count=1 ./api/controllers

test-service:
	go test -count=1 ./api/services
