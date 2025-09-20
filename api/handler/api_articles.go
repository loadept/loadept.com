package handler

import (
	"io"
	"net/http"
	"strings"

	"github.com/loadept/loadept.com/internal/model"
	"github.com/loadept/loadept.com/internal/service"
	"github.com/loadept/loadept.com/pkg/logger"
	"github.com/loadept/loadept.com/pkg/respond"
)

type ApiArticleHandler struct {
	service *service.ArticleService
}

func NewApiArticlesHandler(service *service.ArticleService) *ApiArticleHandler {
	return &ApiArticleHandler{
		service: service,
	}
}

func (h *ApiArticleHandler) GetRawArticleByName(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		respond.JSON(w, respond.Map{
			"detail": "Method '" + r.Method + "' not allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	requestCtx := r.Context()
	category := r.PathValue("category")
	name := r.PathValue("name")

	articles, err := h.service.GetRawArticleByName(requestCtx, category, name)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			respond.JSON(w, respond.Map{
				"detail": "Content not found",
			}, http.StatusNotFound)

			logger.ERROR.Printf("An error occurred while retrieving article: %v\n", err)
			return
		}
		respond.JSON(w, respond.Map{
			"detail": "An error occurred while retrieving results",
		}, http.StatusInternalServerError)

		logger.ERROR.Printf("An error occurred while retrieving article: %v\n", err)
		return
	}
	defer articles.Close()

	if _, err := io.Copy(w, articles); err != nil {
		respond.JSON(w, respond.Map{
			"detail": "An error occurred while retrieving results",
		}, http.StatusInternalServerError)

		logger.ERROR.Printf("An error occurred while generating a response: %v\n", err)
	}
}

func (h *ApiArticleHandler) GetListArticles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		respond.JSON(w, respond.Map{
			"detail": "Method '" + r.Method + "' not allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	requestCtx := r.Context()
	category := r.PathValue("category")

	articles, err := h.service.GetListArticles(requestCtx, category)
	if err != nil {
		if strings.Contains(err.Error(), "404") {
			respond.JSON(w, respond.Map{
				"detail": "Content not found",
			}, http.StatusNotFound)

			logger.ERROR.Printf("An error occurred while listing articles: %v\n", err)
			return
		}
		respond.JSON(w, respond.Map{
			"detail": "An error occurred while retrieving results",
		}, http.StatusInternalServerError)

		logger.ERROR.Printf("An error occurred while listing articles: %v\n", err)
		return
	}

	response := model.ArticleResponse{
		Articles: articles,
	}
	respond.JSON(w, response, http.StatusOK)
}
