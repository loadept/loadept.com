package handler

import (
	"net/http"

	"github.com/loadept/loadept.com/internal/service"
	"github.com/loadept/loadept.com/pkg/logger"
	"github.com/loadept/loadept.com/pkg/respond"
)

type ApiCategoriesHandler struct {
	service *service.CategoryService
}

func NewApiCategoryHandler(service *service.CategoryService) *ApiCategoriesHandler {
	return &ApiCategoriesHandler{
		service: service,
	}
}

func (h *ApiCategoriesHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		respond.JSON(w, respond.Map{
			"detail": "Method '" + r.Method + "' not allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	categories, err := h.service.GetCategories()
	if err != nil {
		respond.JSON(w, respond.Map{
			"detail": "An error occurred while retrieving results",
		}, http.StatusInternalServerError)

		logger.ERROR.Printf("An error occurred while retrieving categories: %v\n", err)
		return
	}

	respond.JSON(w, categories, http.StatusOK)
}
