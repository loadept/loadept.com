package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/loadept/pirca"
	"github.com/loadept/website/internal/middleware"
	"github.com/stretchr/testify/assert"
)

func TestMiddleware_SetsLogEntry(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := pirca.Ctx(r)
		val, _ := ctx.Get("logentry")
		le := val.(*middleware.LogEntry)
		assert.NotNil(t, le)
		assert.Equal(t, "GET", le.Method)
		assert.Equal(t, "/test", le.Path)
		w.WriteHeader(http.StatusOK)
	})
	handler := pirca.New()(middleware.Logger(next))

	req := httptest.NewRequest("GET", "/test", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestMiddleware_CapturesCodes(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
	})
	handler := pirca.New()(middleware.Logger(next))

	req := httptest.NewRequest("GET", "/notfound", nil)
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusNotFound, rec.Code)
}

func TestMiddleware_CaptureCloudflareHeaders(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := pirca.Ctx(r)
		val, _ := ctx.Get("logentry")
		le := val.(*middleware.LogEntry)
		assert.Equal(t, "1.2.3.4", le.Addr)
		assert.Equal(t, "PE", le.Country)
	})
	handler := pirca.New()(middleware.Logger(next))

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("CF-Connecting-IP", "1.2.3.4")
	req.Header.Set("CF-IPCountry", "PE")
	rec := httptest.NewRecorder()

	handler.ServeHTTP(rec, req)
}
