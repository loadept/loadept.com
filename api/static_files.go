package api

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

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
