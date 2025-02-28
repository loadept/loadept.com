package handler

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"
)

func Index() http.HandlerFunc {
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

		tmpl := template.Must(template.ParseFiles("web/template/index.html"))
		tmpl.Execute(w, map[string]int{"Year": currYear})
	}
}

func Docs() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("web/template/docs.html"))

	file, err := os.Open("web/static/data.md")
	if err != nil {
		fmt.Println("Error al leer", err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	buffer := make([]byte, 1250)
	var content []byte

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		content = append(content, buffer[:n]...)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}
