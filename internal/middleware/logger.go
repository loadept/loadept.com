package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"sync"
	"time"
)

var (
	logMu sync.Mutex
	enc   = json.NewEncoder(os.Stdout)
)

type contextKey struct{}

var logEntryKey = contextKey{}

type LogEntry struct {
	Timestamp  string `json:"timestamp"`
	Method     string `json:"method,omitempty"`
	Path       string `json:"path,omitempty"`
	StatusCode int    `json:"status_code,omitempty"`
	CacheHit   bool   `json:"cache_hit,omitempty"`
	Error      string `json:"error,omitempty"`
	Addr       string `json:"ip,omitempty"`
	Country    string `json:"country,omitempty"`
	RayID      string `json:"ray_id,omitempty"`
}

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := &responseRecorder{ResponseWriter: w, statusCode: http.StatusOK}

		le := &LogEntry{
			Timestamp: time.Now().Format(time.RFC3339),
			Method:    r.Method,
			Path:      r.URL.Path,
			Addr:      getClientIP(r),
			Country:   r.Header.Get("CF-IPCountry"),
			RayID:     r.Header.Get("CF-Ray"),
		}
		defer writeLog(le)

		ctx := context.WithValue(r.Context(), logEntryKey, le)
		next.ServeHTTP(rw, r.WithContext(ctx))
		le.StatusCode = rw.statusCode
	})
}

func LogFromCtx(ctx context.Context) *LogEntry {
	le := ctx.Value(logEntryKey).(*LogEntry)
	return le
}

func writeLog(le *LogEntry) {
	logMu.Lock()
	defer logMu.Unlock()
	enc.Encode(le)
}
