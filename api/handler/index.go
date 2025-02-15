package handler

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

type Data struct {
	Content string
}

func Index() http.HandlerFunc {
	tmpl := template.Must(template.ParseFiles("web/template/index.html"))

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
		tmpl.Execute(w, Data{
			Content: string(content),
		})
	}
}
