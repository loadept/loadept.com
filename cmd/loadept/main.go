package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/loadept/loadept.com/api"
	"github.com/loadept/loadept.com/api/handler"
	"github.com/loadept/loadept.com/api/middleware"
)

var (
	addr string
)

func init() {
	PORT := os.Getenv("PORT")
	if len(PORT) == 0 {
		log.Println("The PORT variable is not defined")
		os.Exit(1)
	}
	addr = fmt.Sprintf(":%s", PORT)
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/static/", middleware.GzipEncoding(api.ServeStatic("web/static")))
	mux.Handle("/robots.txt", api.ServeStaticFile("web/static/robots.txt"))
	mux.Handle("/favicon.ico", api.ServeStaticFile("web/static/favicon.ico"))

	mux.Handle("/", handler.Index())

	server := http.Server{
		Addr:    addr,
		Handler: mux,
	}

	log.Printf("\033[32mServer ready to listen on addr %s\033[0m\n", addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("\033[31mError to listen serve\033[0m: %v\n", err)
	}
}
