package v1

import (
	"net/http"

	"github.com/loadept/loadept.com/internal/di"
)

func NewRouter(container *di.Container) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health", container.HealthHandler.Health)
	mux.HandleFunc("GET /categories", container.CategoryHandler.GetCategories)
	mux.HandleFunc("GET /articles/{category}", container.ArticleHandler.GetListArticles)
	mux.HandleFunc("GET /articles/{category}/{name}", container.ArticleHandler.GetRawArticleByName)
	mux.HandleFunc("POST /pdf/compress", container.ApiPDFHandler.GetPDFCompressed)

	return mux
}
