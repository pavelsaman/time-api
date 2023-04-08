ARG API_PORT=8080
FROM golang:latest

WORKDIR /time-api

COPY . .

RUN go mod download

RUN make build

EXPOSE ${API_PORT}

ENV API_VERSION=v1 \
    API_PORT=${API_PORT}

CMD ["./bin/time-api"]