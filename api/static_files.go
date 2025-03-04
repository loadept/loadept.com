package api

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// ServeStatic defines a controller to serve static files from a directory.
//
// It is strictly necessary that the endpoint definition has the name /static/
//
// This way, uniqueness is guaranteed for accessing static files.
func ServeStatic(staticDir string) http.Handler {
	fs := http.FileServer(http.Dir(staticDir))

	return http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cleanedUrl := filepath.Clean(r.URL.Path)
		path := filepath.Join(staticDir, cleanedUrl)

		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			http.Error(w, "This method is not allowed", http.StatusMethodNotAllowed)
			return
		}
		if strings.HasPrefix(filepath.Base(r.URL.Path), ".") {
			http.Error(w, "404 page not found", http.StatusNotFound)
			return
		}

		info, err := os.Stat(path)
		if err != nil {
			fs.ServeHTTP(w, r)
			return
		}
		if info.IsDir() {
			http.Error(w, "404 page not found", http.StatusNotFound)
			return
		}

		fs.ServeHTTP(w, r)
	}))
}

// ServeDir defines a controller to serve static files from a directory.
func ServeSPA(staticDir, indexFile string) http.Handler {
	fs := http.FileServer(http.Dir(staticDir))

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cleanedUrl := filepath.Clean(r.URL.Path)
		path := filepath.Join(staticDir, cleanedUrl)

		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			http.Error(w, "This method is not allowed", http.StatusMethodNotAllowed)
			return
		}
		if strings.HasPrefix(filepath.Base(r.URL.Path), ".") {
			http.Error(w, "404 page not found", http.StatusNotFound)
			return
		}

		info, err := os.Stat(path)
		if os.IsNotExist(err) {
			http.ServeFile(w, r, filepath.Join(staticDir, indexFile))
			return
		} else if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		if info.IsDir() && r.URL.Path != "/" {
			http.ServeFile(w, r, filepath.Join(staticDir, indexFile))
			return
		}

		fs.ServeHTTP(w, r)
	})
}

// ServeStaticFile defines a handler that returns a static file from a data path.
//
// The name of the endpoint is not strict, as in the case of "ServeStatic"
//
// which must have a /static/ prefix.
func ServeStaticFile(staticFile string) http.Handler {
	baseDir := filepath.Dir(staticFile)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			http.Error(w, "This method is not allowed", http.StatusMethodNotAllowed)
			return
		}

		requestedPath := r.URL.Path
		cleanedPath := filepath.Clean(requestedPath)
		fullPath := filepath.Join(baseDir, cleanedPath)

		if !strings.HasPrefix(fullPath, baseDir) {
			http.Error(w, "404 page not found", http.StatusNotFound)
			return
		}
		if strings.HasPrefix(filepath.Base(r.URL.Path), ".") {
			http.Error(w, "404 page not found", http.StatusNotFound)
			return
		}
		if cleanedPath != requestedPath {
			http.Error(w, "404 page not found", http.StatusNotFound)
			return
		}
		http.ServeFile(w, r, fullPath)
	})
}
