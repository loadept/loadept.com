package main

import (
	"fmt"
	"net/http"

	"github.com/loadept/loadept.com/api"
	"github.com/loadept/loadept.com/api/handler"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", api.ServeStatic("web/static"))
	mux.Handle("/robots.txt", api.ServeStaticFile("web/static/robots.txt"))

	mux.Handle("/", handler.Index())

	mux.Handle("/home", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hola mundo</h1>"))
	}))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("\033[32mServer ready to listen\033[0m")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("\033[31mError to listen serve\033[0m: %v", err)
	}
}
