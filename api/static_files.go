package api

import (
	"net/http"
	"path/filepath"
	"strings"
)

func ServeStatic(staticDir string) http.Handler {
	fs := http.FileServer(http.Dir(staticDir))

	return http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "This method is not allowed", http.StatusMethodNotAllowed)
			return
		}
		if strings.HasPrefix(filepath.Base(r.URL.Path), ".") {
			http.Error(w, "404 page not found", http.StatusNotFound)
			return
		}
		fs.ServeHTTP(w, r)
	}))
}
