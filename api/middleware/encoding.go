package middleware

import (
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

// gzipResponseWriter is a structure that extends the ResponseWriter interface,
//
// becoming a custom ResponseWriter.
type gzipResponseWriter struct {
	http.ResponseWriter
	Writer io.Writer
}

// Write overrides the write method by changing the default writer to the gzip writer.
func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

// GzipEncoding is a middleware that compresses the http response using the
//
//	compress/gzip
//
// package, wrapping the response if the client supports it, otherwise doing nothing.
func GzipEncoding(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			next.ServeHTTP(w, r)
			return
		}

		w.Header().Set("Content-Encoding", "gzip")
		w.Header().Set("Vary", "Accept-Encoding")

		gz := gzip.NewWriter(w)
		defer gz.Close()

		gw := gzipResponseWriter{ResponseWriter: w, Writer: gz}
		next.ServeHTTP(gw, r)
	}
}
