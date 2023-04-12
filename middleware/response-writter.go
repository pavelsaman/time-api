package middleware

import (
	"net/http"
	"strings"
)

type responseWriter struct {
	http.ResponseWriter
	status int
	body   []byte
}

func (rl *responseWriter) WriteHeader(status int) {
	rl.status = status
	rl.ResponseWriter.WriteHeader(status)
}

func (rw *responseWriter) Write(data []byte) (int, error) {
	rw.body = data
	rw.body = []byte(strings.ReplaceAll(string(rw.body), "\n", ""))
	return rw.ResponseWriter.Write(data)
}
