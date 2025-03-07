package handler

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/loadept/loadept.com/internal/model"
	"github.com/loadept/loadept.com/internal/service"
	"github.com/loadept/loadept.com/pkg/respond"
	"github.com/loadept/loadept.com/pkg/util"
)

type ApiCategoriesHandler struct {
	service *service.CategoryService
}

func NewApiCategoryHandler(service *service.CategoryService) *ApiCategoriesHandler {
	return &ApiCategoriesHandler{
		service: service,
	}
}

func (h *ApiCategoriesHandler) RegisterCategory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respond.JSON(w, respond.Map{
			"detail": "Method '" + r.Method + "' not allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	var category model.CategoryModel
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		var errorDetail string
		switch {
		case errors.Is(err, io.EOF):
			errorDetail = "Request body is empty"
		case strings.Contains(err.Error(), "cannot unmarshal"):
			errorDetail = "Invalid data type in request"
		case strings.Contains(err.Error(), "invalid character"):
			errorDetail = "Invalid JSON format"
		default:
			errorDetail = "Error processing request body"
		}
		respond.JSON(w, respond.Map{
			"detail": errorDetail,
		}, http.StatusBadRequest)
		return
	}

	err := h.service.RegisterCategory(&category)
	if err != nil {
		if _, ok := err.(util.ValidationError); ok {
			respond.JSON(w, respond.Map{
				"detail": err.Error(),
			}, http.StatusBadRequest)
			return
		}
		respond.JSON(w, respond.Map{
			"detail": "An error occurred while registering the category",
		}, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Location", fmt.Sprintf("/api/category/%s", category.ID))
	respond.JSON(w, respond.Map{
		"message":     "Category inserted successfully",
		"category_id": category.ID,
	}, http.StatusCreated)
}

func (h *ApiCategoriesHandler) GetCategoryByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		respond.JSON(w, respond.Map{
			"detail": "Method '" + r.Method + "' not allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	id := r.PathValue("id")

	category, err := h.service.GetCategoryByID(id)
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

	respond.JSON(w, category, http.StatusOK)
}

func (h *ApiCategoriesHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		respond.JSON(w, respond.Map{
			"detail": "Method '" + r.Method + "' not allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	page := r.URL.Query().Get("page")
	if len(page) == 0 {
		page = "1"
	}

	categories, err := h.service.GetCategories(page)
	if err != nil {
		respond.JSON(w, respond.Map{
			"detail": "An error occurred while retrieving results",
		}, http.StatusInternalServerError)
		return
	}

	count := strconv.Itoa(len(categories))
	response := model.CategoryResponse{
		Page:  page,
		Count: count,
		Data:  categories,
	}
	respond.JSON(w, response, http.StatusOK)
}
