package middleware

import (
	"io"
	"net/http"
	"strings"

	"github.com/andybalholm/brotli"
)

// brotliResponseWriter is a structure that extends the ResponseWriter interface,
//
// becoming a custom ResponseWriter.
type brotliResponseWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

// Write overrides the write method by changing the default writer to the brotli writer.
func (w brotliResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// BrotliEncorder is a middleware that compresses the http response using the
//
//	github.com/andybalholm/brotli
//
// package, wrapping the response if the client supports it, otherwise doing nothing.
func BrotliEncorder(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "br") {
			next.ServeHTTP(w, r)
			return
		}

		w.Header().Set("Content-Encoding", "br")
		w.Header().Set("Vary", "Accept-Encoding")

		brl := brotli.NewWriter(w)
		defer brl.Close()

		bw := brotliResponseWriter{ResponseWriter: w, Writer: brl}
		next.ServeHTTP(bw, r)
	})
}
