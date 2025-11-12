package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func mockHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
}

func TestGzipEncodig(t *testing.T) {
	t.Run("No gzip require", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		rr := httptest.NewRecorder()

		handler := GzipEncoding(mockHandler())
		handler.ServeHTTP(rr, req)

		if rr.Header().Get("Content-Encoding") == "gzip" {
			t.Errorf("Unexpected gzip encoding")
		}

		expectedBody := "Hello, World!"
		if rr.Body.String() != expectedBody {
			t.Errorf("Expected body %q, got %q", expectedBody, rr.Body.String())
		}
	})

	t.Run("Gzip require", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/", nil)
		req.Header.Set("Accept-Encoding", "gzip")
		rr := httptest.NewRecorder()

		handler := GzipEncoding(mockHandler())
		handler.ServeHTTP(rr, req)

		if rr.Header().Get("Content-Encoding") != "gzip" {
			t.Errorf("Expected gzip encoding, got %q", rr.Header().Get("Content-Encoding"))
		}

		reader, err := gzip.NewReader(rr.Body)
		if err != nil {
			t.Fatalf("Failed to create gzip reader: %v", err)
		}
		defer reader.Close()

		uncompressedBody, err := io.ReadAll(reader)
		if err != nil {
			t.Fatalf("Failed to read gzip response: %v", err)
		}

		expectedBody := "Hello, World!"
		if string(uncompressedBody) != expectedBody {
			t.Errorf("Expected body %q, got %q", expectedBody, string(uncompressedBody))
		}
	})
}
