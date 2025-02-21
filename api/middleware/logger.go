package middleware

import (
	"log"
	"net/http"
	"time"
)

type loggerResponseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (lrw *loggerResponseWriter) WriteHeader(code int) {
	lrw.statusCode = code
	lrw.ResponseWriter.WriteHeader(code)
}

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
		log.Printf("\033[33m-\033[0m %s %s: %s %s %s - %d %s\n",
			forwardedProto, forwardedHost, forwardedIp,
			method, requestedPath, lwr.statusCode, elapsed,
		)
	})
}
