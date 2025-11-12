package validation

import (
	"net/http"
	"strings"

	httpError "github.com/loadept/loadept.com/pkg/respond/error"
)

func ValidateUploadedFileRequest(r *http.Request) error {
	contentType := r.Header.Get("Content-Type")
	if !strings.HasPrefix(contentType, "multipart/form-data") {
		return &httpError.FieldError{
			Field:   "file",
			Message: "Request must be multipart/form-data",
		}
	}
	if r.ContentLength == 0 {
		return &httpError.FieldError{
			Field:   "file",
			Message: "Empty body",
		}
	}
	return nil
}
