package handler

import (
	"database/sql"
	"net/http"
	"strconv"

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

func (h *ApiArticleHandler) GetArticleByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		respond.JSON(w, respond.Map{
			"detail": "Method '" + r.Method + "' not allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")

	article, err := h.service.GetArticleByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			respond.JSON(w, respond.Map{
				"detail": "No results found",
			}, http.StatusNotFound)
			return
		}
		respond.JSON(w, respond.Map{
			"detail": "An error occurred while retrieving results",
		}, http.StatusInternalServerError)
		return
	}

	respond.JSON(w, article, http.StatusOK)
}

func (h *ApiArticleHandler) GetRecentArticles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		respond.JSON(w, respond.Map{
			"detail": "Method '" + r.Method + "' not allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	category := r.URL.Query().Get("category")

	articles, err := h.service.GetRecentArticles(category)
	if err != nil {
		respond.JSON(w, respond.Map{
			"detail": "An error occurred while retrieving results",
		}, http.StatusInternalServerError)
		return
	}

	count := strconv.Itoa(len(articles))
	response := model.ArticleResponse{
		Count: count,
		Data:  articles,
	}
	respond.JSON(w, response, http.StatusOK)
}

func (h *ApiArticleHandler) GetArticles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		respond.JSON(w, respond.Map{
			"detail": "Method '" + r.Method + "' not allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	category := r.URL.Query().Get("category")
	title := r.URL.Query().Get("search")
	page := r.URL.Query().Get("page")
	if len(page) == 0 {
		page = "1"
	}

	articles, err := h.service.GetArticles(category, title, page)
	if err != nil {
		respond.JSON(w, respond.Map{
			"detail": "An error occurred while retrieving results",
		}, http.StatusInternalServerError)
		return
	}

	count := strconv.Itoa(len(articles))
	response := model.ArticleResponse{
		Page:  page,
		Count: count,
		Data:  articles,
	}
	respond.JSON(w, response, http.StatusOK)
}
