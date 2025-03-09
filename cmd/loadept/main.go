package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/validator/v10"
	"github.com/loadept/loadept.com/api"
	"github.com/loadept/loadept.com/api/handler"
	"github.com/loadept/loadept.com/api/middleware"
	"github.com/loadept/loadept.com/internal/auth"
	"github.com/loadept/loadept.com/internal/config"
	"github.com/loadept/loadept.com/internal/database"
	"github.com/loadept/loadept.com/internal/repository"
	"github.com/loadept/loadept.com/internal/service"
)

var conn *sql.DB

func init() {
	config.LoadConfig()

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

func main() {
	mux := http.NewServeMux()

	validate := validator.New()

	authService := auth.NewAuthService("JWT")
	authMiddleware := middleware.NewAuthMiddleware(authService)

	userRepo := repository.NewUserRepository(conn)
	userService := service.NewUserService(userRepo, validate)
	userHandler := handler.NewApiUserHandler(userService, authService)

	articleRepo := repository.NewArticleRepository(conn)
	articleService := service.NewArticleService(articleRepo, validate)
	handlerArticles := handler.NewArticlesHandler(articleService)

	categoryRepo := repository.NewCetogoryRepository(conn)
	categoryService := service.NewCategoryService(categoryRepo, validate)
	categoryHandler := handler.NewApiCategoryHandler(categoryService)

	// File Server
	mux.Handle("/", middleware.GzipEncoding(api.ServeSPA("web/dist", "index.html")))

	// Accessible routes
	{
		mux.HandleFunc("/api/category/{id}", categoryHandler.GetCategoryByID)
		mux.HandleFunc("/api/category", categoryHandler.GetCategories)

		mux.HandleFunc("/api/article/recent", handlerArticles.GetRecentArticles)
		mux.HandleFunc("/api/article/{id}", handlerArticles.GetArticleByID)
		mux.HandleFunc("/api/article", handlerArticles.GetArticles)

		mux.HandleFunc("/api/auth/register", userHandler.RegisterUser)
		mux.HandleFunc("/api/auth/login", userHandler.LoginUser)
	}

	// Protected router for admin
	{
		mux.Handle("/api/category/register", authMiddleware(
			http.HandlerFunc(categoryHandler.RegisterCategory),
		))
		mux.Handle("/api/article/register", authMiddleware(
			http.HandlerFunc(handlerArticles.RegisterArticle),
		))
	}

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
