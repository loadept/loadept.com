package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/loadept/loadept.com/api"
	"github.com/loadept/loadept.com/api/middleware"
	"github.com/loadept/loadept.com/internal/config"
)

func init() {
	config.LoadConfig()
}

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", middleware.GzipEncoding(api.ServeSPA("web/dist")))

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", config.Env.PORT),
		Handler: middleware.LoggerMiddleware(mux),
	}

	log.Printf("\033[32mServer ready to listen on addr %s\033[0m\n", config.Env.PORT)
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("\033[31mError to listen serve\033[0m: %v\n", err)
	}
}
