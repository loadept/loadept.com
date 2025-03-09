package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/loadept/loadept.com/internal/auth"
	"github.com/loadept/loadept.com/internal/model"
	"github.com/loadept/loadept.com/internal/service"
	"github.com/loadept/loadept.com/pkg/respond"
	"github.com/loadept/loadept.com/pkg/util"
)

type ApiUserHandler struct {
	service      *service.UserService
	tokenService auth.TokenService
}

func NewApiUserHandler(service *service.UserService, tokenService auth.TokenService) *ApiUserHandler {
	return &ApiUserHandler{
		service:      service,
		tokenService: tokenService,
	}
}

func (h *ApiUserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respond.JSON(w, respond.Map{
			"detail": "Method '" + r.Method + "' not allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	var user model.UserModel
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
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

	result, err := h.service.GetUser(&user)
	if err != nil {
		if _, ok := err.(util.ValidationError); ok {
			respond.JSON(w, respond.Map{
				"detail": err.Error(),
			}, http.StatusBadRequest)
			return
		} else if strings.Contains(err.Error(), "Incorrect") {
			respond.JSON(w, respond.Map{
				"detail": err.Error(),
			}, http.StatusBadRequest)
			return
		}
		respond.JSON(w, respond.Map{
			"detail": "An error occurred while registering user",
		}, http.StatusInternalServerError)
		return
	}

	token, err := h.tokenService.CreateToken(result.ID, result.IsAdmin)
	if err != nil {
		respond.JSON(w, respond.Map{
			"detail": "An error occurred while creating auth token",
		}, http.StatusInternalServerError)
	}

	respond.JSON(w, respond.Map{
		"auth_token": token,
	}, http.StatusCreated)
}

func (h *ApiUserHandler) RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respond.JSON(w, respond.Map{
			"detail": "Method '" + r.Method + "' not allowed",
		}, http.StatusMethodNotAllowed)
		return
	}

	var user model.RegisterUserModel
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
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

	err := h.service.RegisterUser(&user)
	if err != nil {
		if _, ok := err.(util.ValidationError); ok {
			respond.JSON(w, respond.Map{
				"detail": err.Error(),
			}, http.StatusBadRequest)
			return
		} else if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			respond.JSON(w, respond.Map{
				"detail": fmt.Sprintf("User '%s/%s' already exists on this system",
					user.Username,
					user.Email,
				),
			}, http.StatusBadRequest)
			return
		}
		respond.JSON(w, respond.Map{
			"detail": "An error occurred while registering user",
		}, http.StatusInternalServerError)
		return
	}

	respond.JSON(w, respond.Map{
		"message": "User inserted successfully",
		"user_id": user.ID,
	}, http.StatusCreated)
}
