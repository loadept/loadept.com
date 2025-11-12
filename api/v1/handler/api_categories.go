package handler

import (
	"net/http"

	application "github.com/loadept/loadept.com/internal/application/category"
	"github.com/loadept/loadept.com/pkg/logger"
	"github.com/loadept/loadept.com/pkg/respond"
)

type ApiCategoriesHandler struct {
	service *application.CategoryService
}

func NewApiCategoryHandler(service *application.CategoryService) *ApiCategoriesHandler {
	return &ApiCategoriesHandler{
		service: service,
	}
}

func (h *ApiCategoriesHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	requestCtx := r.Context()

	categories, err := h.service.GetCategories(requestCtx)
	if err != nil {
		respond.JSON(w, respond.Map{
			"detail": "An error occurred while retrieving results",
		}, http.StatusInternalServerError)

		logger.ERROR.Printf("An error occurred while retrieving categories: %v\n", err)
		return
	}

	respond.JSON(w, categories, http.StatusOK)
}
