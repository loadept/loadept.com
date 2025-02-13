package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/home", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hola mundo</h1>"))
	}))

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	fmt.Println("Serve ready to listen")
	err := server.ListenAndServe()
	if err != nil {
		fmt.Printf("Error to listen serve: %v", err)
	}
}
