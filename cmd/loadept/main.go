package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/loadept/loadept.com/api/middleware"
	apiv1 "github.com/loadept/loadept.com/api/v1"
	"github.com/loadept/loadept.com/internal/config"
	"github.com/loadept/loadept.com/internal/di"
	"github.com/loadept/loadept.com/internal/infrastructure/cache"
	"github.com/loadept/loadept.com/internal/infrastructure/database"
	"github.com/loadept/loadept.com/pkg/logger"
)

var (
	ctx = context.Background()
)

func init() {
	config.LoadEnviron()
	config.LoadConfig()
	logger.NewLogger()
}

func main() {
	defer logger.CloseLogger()
	mux := http.NewServeMux()
	httpClient := &http.Client{}
	pdfServiceUrl, _ := url.Parse(config.Env.PDF_SERVICE_URL)

	db, err := database.NewConnection()
	if err != nil {
		log.Println("Error to connect with database:", err)
	}
	rdb, err := cache.NewRedisConnection(ctx)
	if err != nil {
		log.Println("Error to connect with redis:", err)
	}

	container := di.NewContainer(
		db.GetDB(),
		rdb.GetClient(),
		httpClient,
		httputil.NewSingleHostReverseProxy(pdfServiceUrl),
	)
	container.Build(ctx)

	apiRouter := apiv1.NewRouter(container)
	mux.Handle("/api/v1/", http.StripPrefix("/api/v1", apiRouter))
	mux.HandleFunc("GET /sitemap.xml", container.SitemapHandler.GetSitemap)

	var handler http.Handler = mux
	handler = middleware.LoggerMiddleware(handler)
	if config.Env.DEBUG == "true" {
		log.Println("You are in debug mode, cors middleware will be used")
		handler = middleware.CorsMiddleware(handler)
	}

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", config.Env.PORT),
		Handler: handler,
	}
	log.Printf("Server ready to listen on addr %s", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Error to listen serve: %v\n", err)
	}
}
