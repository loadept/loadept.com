package handler

import (
	"bytes"
	"errors"
	"io"
	"net/http"
	"net/http/httputil"

	"github.com/loadept/loadept.com/internal/validation"
	"github.com/loadept/loadept.com/pkg/logger"
	"github.com/loadept/loadept.com/pkg/respond"
	httpError "github.com/loadept/loadept.com/pkg/respond/error"
)

type ApiPDFHandler struct {
	proxy *httputil.ReverseProxy
}

func NewApiPDFHandler(proxy *httputil.ReverseProxy) *ApiPDFHandler {
	return &ApiPDFHandler{
		proxy: proxy,
	}
}

func (h *ApiPDFHandler) GetPDFCompressed(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		respond.JSON(w, respond.Map{
			"detail": "An error ocurred while validate data",
		}, http.StatusInternalServerError)
		return
	}
	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	if err := validation.ValidateUploadedFileRequest(r); err != nil {
		var validationErr *httpError.FieldError
		if errors.As(err, &validationErr) {
			respond.JSON(w, respond.Map{
				"detail": "Validation failed",
				"error":  validationErr,
			}, http.StatusBadRequest)
			return
		}
		respond.JSON(w, respond.Map{
			"detail": "An error ocurred while validate data",
		}, http.StatusInternalServerError)
		logger.ERROR.Printf("Validation error: %v\n", err)
		return
	}

	quality := r.FormValue("quality")
	if len(quality) < 1 {
		respond.JSON(w, respond.Map{
			"detail": "Invalid request",
			"error": &httpError.FieldError{
				Field:   "quality",
				Message: "Quality field is required",
			},
		}, http.StatusBadRequest)
		return
	}

	r.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	h.proxy.ServeHTTP(w, r)
}
