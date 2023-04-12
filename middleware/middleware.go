package middleware

import (
	"net/http"
	"os"

	console_logger "github.com/pavelsaman/time-api/middleware/logger/loggers"
)

var consoleLogger console_logger.ConsoleLogger = console_logger.ConsoleLogger{}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		consoleLogger.Info(os.Stdout, "Received %s request for %s from %s\n", r.Method, r.URL.Path, r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func LogResponses(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseWriter{ResponseWriter: w, status: http.StatusOK}
		defer func() {
			logLevel := consoleLogger.Info
			writer := os.Stdout
			if rw.status >= 400 && rw.status < 500 {
				logLevel = consoleLogger.Warn
			} else if rw.status >= 500 {
				writer = os.Stderr
				logLevel = consoleLogger.Error
			}

			logLevel(writer, "Response for %s request for %s from %s: %d %s\n", r.Method, r.URL.Path, r.RemoteAddr, rw.status, rw.body)
		}()
		next.ServeHTTP(rw, r)
	})
}
