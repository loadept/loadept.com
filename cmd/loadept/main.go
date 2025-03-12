package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/loadept/loadept.com/api"
	"github.com/loadept/loadept.com/api/handler"
	"github.com/loadept/loadept.com/api/middleware"
	"github.com/loadept/loadept.com/internal/config"
	"github.com/loadept/loadept.com/internal/infrastructure/cache"
	"github.com/loadept/loadept.com/internal/infrastructure/database"
	"github.com/loadept/loadept.com/internal/service"
	"github.com/redis/go-redis/v9"
)

var (
	conn *sql.DB
	rdb  *redis.Client
	ctx  = context.Background()
)

func init() {
	config.LoadConfig()

	{
		db, err := database.NewConnection()
		if err != nil {
			log.Println("Error to connect with database")
			os.Exit(1)
		}

		now, err := db.GetNow()
		if err != nil {
			log.Println("Error to optain current date")
			os.Exit(1)
		}

		formatDate := now.Format("2006-01-02")
		log.Printf("Connection with database success, current date %s\n", formatDate)

		conn = db.GetDB()
	}
	{
		redisClient, err := cache.NewRedisConnection(ctx)
		if err != nil {
			log.Println("Error to connect with redis:", err)
			os.Exit(1)
		}

		redisDate, err := redisClient.GetNow()
		if err != nil {
			log.Println("Error to optain current date from redis:", err)
			os.Exit(1)
		}

		formatRedisDate := redisDate.Format("2006-01-02")
		log.Printf("Connection with redis success, current date %s\n", formatRedisDate)

		rdb = redisClient.GetClient()
	}
}

func main() {
	mux := http.NewServeMux()
	httpClient := &http.Client{}

	articleService := service.NewArticleService(httpClient, rdb, ctx)
	handlerArticles := handler.NewArticlesHandler(articleService)

	categoryService := service.NewCategoryService(httpClient, rdb, ctx)
	categoryHandler := handler.NewApiCategoryHandler(categoryService)

	{
		mux.Handle("/", middleware.GzipEncoding(api.ServeSPA("web/dist", "index.html")))

		mux.HandleFunc("/api/category", categoryHandler.GetCategories)
		mux.HandleFunc("/api/article/{category}", handlerArticles.GetArticles)
		mux.HandleFunc("/api/article/{category}/{name}", handlerArticles.GetArticleBySha)
		mux.HandleFunc("/api/article/edit/{category}/{name}", handlerArticles.EditArticle)
	}

	debug, err := strconv.ParseBool(config.Env.DEBUG)
	if err != nil {
		log.Printf("Error to parse to bool: %v", err)
	}

	var muxWrapped http.Handler
	if debug {
		log.Println("You are in \033[33mdebug\033[0m mode, cors middleware will be used")
		muxWrapped = middleware.CorsMiddleware(
			middleware.LoggerMiddleware(mux),
		)
	} else {
		muxWrapped = middleware.LoggerMiddleware(mux)
	}

	server := http.Server{
		Addr:    fmt.Sprintf(":%s", config.Env.PORT),
		Handler: muxWrapped,
	}

	log.Printf("\033[32mServer ready to listen on addr %s\033[0m\n", server.Addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("\033[31mError to listen serve\033[0m: %v\n", err)
	}
}
