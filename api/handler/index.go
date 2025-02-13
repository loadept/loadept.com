package handler

import (
	"bufio"
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/template/index.html")
	if err != nil {
		fmt.Println(err)
	}

	file, err := os.Open("data.md")
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

	tmpl.ExecuteTemplate(w, "index.html", map[string]string{
		"md": string(content),
	})
}
