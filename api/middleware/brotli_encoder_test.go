package middleware

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/andybalholm/brotli"
	"github.com/stretchr/testify/assert"
)

func mockBrotliHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
}

func TestBrotliEncoder(t *testing.T) {
	t.Run("No brotli require", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		rr := httptest.NewRecorder()

		handler := BrotliEncorder(mockBrotliHandler())
		handler.ServeHTTP(rr, req)

		contentHeader := rr.Header().Get("Content-Encoding")
		assert.NotEqual(t, "br", contentHeader, "Unexpected brotli encoding")

		expectedBody := "Hello, World!"
		assert.Equal(t, expectedBody, rr.Body.String(), "The response body does not match")
	})

	t.Run("Brotli require", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Accept-Encoding", "br")
		rr := httptest.NewRecorder()

		handler := BrotliEncorder(mockBrotliHandler())
		handler.ServeHTTP(rr, req)

		contentHeader := rr.Header().Get("Content-Encoding")
		assert.Equal(t, "br", contentHeader, "Unexpected brotli encoding")

		reader := brotli.NewReader(rr.Body)

		uncompressedBody, err := io.ReadAll(reader)
		assert.Nil(t, err, "Failed to read brotli response")

		expectedBody := "Hello, World!"
		assert.Equal(t, expectedBody, string(uncompressedBody), "The response body does not match")
	})
}
