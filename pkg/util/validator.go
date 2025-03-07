package util

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

type ValidationError struct {
	message string
}

func (e ValidationError) Error() string {
	return e.message
}

func HandleValidationErrors(err error) error {
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, e := range validationErrors {
			field := strings.ToLower(e.Field())
			var message string

			switch e.Tag() {
			case "required":
				message = fmt.Sprintf("%s field is required", field)
			case "min":
				message = fmt.Sprintf("Minimum length of the %s field is %s", field, e.Param())
			case "max":
				message = fmt.Sprintf("Maximum length of the %s field is %s", field, e.Param())
			case "email":
				message = fmt.Sprintf("%s field must be a valid email", field)
			default:
				message = fmt.Sprintf("Validation error: %s", e.Tag())
			}
			return ValidationError{message: message}
		}
	}
	return err
}
