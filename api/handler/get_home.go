package handler

import (
	"html/template"
	"net/http"
	"time"
)

func Index() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles(
		"web/template/layout/base.tmpl",
		"web/template/layout/header.tmpl",
		"web/template/layout/footer.tmpl",
		"web/template/index.html",
	))

	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			http.Error(w, "This method is not allowed", http.StatusMethodNotAllowed)
			return
		}

		currYear := time.Now().Year()

		content := struct {
			Year int
		}{
			Year: currYear,
		}

		tmpl.ExecuteTemplate(w, "base", content)
	}
}
