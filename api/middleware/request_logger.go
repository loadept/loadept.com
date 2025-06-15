package middleware

import (
	"net/http"
	"time"

	"github.com/loadept/loadept.com/pkg/logger"
)

// loggerResponseWriter is a structure that extends the ResponseWriter interface,
//
// adding the statusCode field becoming a custom ResponseWriter.
type loggerResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader overrides the Writeheader method, assigning the status code to
// the new
//
//	statusCode
//
// field
func (lrw *loggerResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

// LoggerMiddleware writes logs of requests forwarded by the proxy to tty,
//
// as well as taking the statusCode field with the value assigned in WriteHeader.
func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		lwr := &loggerResponseWriter{ResponseWriter: w, statusCode: http.StatusOK}

		method := r.Method
		requestedPath := r.URL.Path
		forwardedIp := r.Header.Get("X-Forwarded-For")
		forwardedHost := r.Header.Get("X-Forwarded-Host")
		forwardedProto := r.Header.Get("X-Forwarded-Proto")

		next.ServeHTTP(lwr, r)

		elapsed := time.Since(start)
		logger.INFO.Printf("%s %s: %s %s %s - %d %s\n",
			forwardedProto, forwardedHost, forwardedIp,
			method, requestedPath, lwr.statusCode, elapsed,
		)
	})
}
