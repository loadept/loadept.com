package middleware_test

import (
	"bytes"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/loadept/website/internal/middleware"
	"github.com/stretchr/testify/assert"
)

func captureLogs(t *testing.T) *bytes.Buffer {
	t.Helper()
	var buf bytes.Buffer
	prev := slog.Default()
	slog.SetDefault(slog.New(slog.NewTextHandler(&buf, nil)))
	t.Cleanup(func() { slog.SetDefault(prev) })
	return &buf
}

func TestLogger_LogsRequestFields(t *testing.T) {
	buf := captureLogs(t)
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	handler := middleware.Logger(next)

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	log := buf.String()
	assert.Contains(t, log, "request completed")
	assert.Contains(t, log, "method=GET")
	assert.Contains(t, log, "path=/test")
	assert.Contains(t, log, "status=200")
}

func TestLogger_CapturesNonDefaultStatus(t *testing.T) {
	buf := captureLogs(t)
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	handler := middleware.Logger(next)

	req := httptest.NewRequest("GET", "/notfound", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusNotFound, rec.Code)
	assert.Contains(t, buf.String(), "status=404")
}

func TestLogger_CaptureCloudflareHeaders(t *testing.T) {
	buf := captureLogs(t)
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	handler := middleware.Logger(next)

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("CF-Connecting-IP", "1.2.3.4")
	req.Header.Set("CF-IPCountry", "PE")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)

	log := buf.String()
	assert.Contains(t, log, "addr=1.2.3.4")
	assert.Contains(t, log, "country=PE")
}
