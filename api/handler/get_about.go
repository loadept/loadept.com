package handler

import (
	"html/template"
	"net/http"
	"time"
)

func About() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles(
		"web/template/layout/base.tmpl",
		"web/template/layout/header.tmpl",
		"web/template/layout/footer.tmpl",
		"web/template/about.html",
	))

	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet && r.Method != http.MethodHead {
			http.Error(w, "This method is not allowed", http.StatusMethodNotAllowed)
			return
		}

		currYear := time.Now().Year()

		tmpl.ExecuteTemplate(w, "base", map[string]int{"Year": currYear})
	}
}
