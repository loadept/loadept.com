package api

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func ServeStatic(staticDir string) http.Handler {
	fs := http.FileServer(http.Dir(staticDir))

	return http.StripPrefix("/static/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cleanedUrl := filepath.Clean(r.URL.Path)
		path := filepath.Join(staticDir, cleanedUrl)

		if r.Method != http.MethodGet {
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
			fmt.Println("es dir bro ", info.Name())
			return
		}

		fs.ServeHTTP(w, r)
	}))
}

func ServeStaticFile(staticFile string) http.Handler {
	baseDir := filepath.Dir(staticFile)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "This method is not allowed", http.StatusMethodNotAllowed)
			return
		}

		requestedPath := r.URL.Path
		cleanedPath := filepath.Clean(requestedPath)
		fullPath := filepath.Join(baseDir, cleanedPath)

		if !strings.HasPrefix(fullPath, baseDir) {
			fmt.Println("Entra")
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
