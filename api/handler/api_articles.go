package handler

import (
	"io"
	"net/http"
	"strings"

	"github.com/loadept/loadept.com/internal/model"
	"github.com/loadept/loadept.com/internal/service"
	"github.com/loadept/loadept.com/pkg/respond"
)

type ApiArticleHandler struct {
	service *service.ArticleService
}

func NewArticlesHandler(service *service.ArticleService) *ApiArticleHandler {
	return &ApiArticleHandler{
		service: service,
	}
}

// Pseudo cache
var cache = make(map[string][]model.ArticleModel)

func (h *ApiArticleHandler) GetArticleBySha(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		respond.JSON(w, respond.Map{
			"detail": "Method '" + r.Method + "' not allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	category := r.PathValue("category")
	name := r.PathValue("name")

	articles, err := h.service.GetArticleBySha(category, name)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			respond.JSON(w, respond.Map{
				"detail": "Content not found",
			}, http.StatusNotFound)
			return
		}
		respond.JSON(w, respond.Map{
			"detail": "An error occurred while retrieving results",
		}, http.StatusInternalServerError)
		return
	}
	defer articles.Close()

	if _, err := io.Copy(w, articles); err != nil {
		respond.JSON(w, respond.Map{
			"detail": "An error occurred while retrieving results",
		}, http.StatusInternalServerError)
	}
}

func (h *ApiArticleHandler) GetArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		respond.JSON(w, respond.Map{
			"detail": "Method '" + r.Method + "' not allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	category := r.PathValue("category")

	if artc, ok := cache[category]; ok {
		respond.JSON(w, respond.Map{
			"articles": artc,
		}, http.StatusOK)
		return
	}

	articles, err := h.service.GetArticles(category)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			respond.JSON(w, respond.Map{
				"detail": "Content not found",
			}, http.StatusNotFound)
			return
		}
		respond.JSON(w, respond.Map{
			"detail": "An error occurred while retrieving results",
		}, http.StatusInternalServerError)
		return
	}
	cache[category] = articles

	response := model.ArticleResponse{
		Articles: articles,
	}
	respond.JSON(w, response, http.StatusOK)
}

func (h *ApiArticleHandler) EditArticle(w http.ResponseWriter, r *http.Request) {
	category := r.PathValue("category")
	name := r.PathValue("name")

	url := h.service.EditArticle(category, name)
	http.Redirect(w, r, url, http.StatusFound)
}
