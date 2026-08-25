package middleware

import (
	"log/slog"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		rec := &statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(rec, r)

		slog.Info(
			"request completed",
			"method", r.Method,
			"path", r.URL.Path,
			"addr", getClientIP(r),
			"country", r.Header.Get("CF-IPCountry"),
			"status", rec.status,
			"duration_ms", time.Since(start).Microseconds(),
		)
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}
