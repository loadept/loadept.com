package util

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name  string `validate:"required,max=10,min=3"`
	Email string `validate:"required,email"`
}

func TestHandleValidationErrors(t *testing.T) {
	validate := validator.New()

	t.Run("Fields not available", func(t *testing.T) {
		testData := TestStruct{}
		err := validate.Struct(testData)
		if err != nil {
			handlerErr := HandleValidationErrors(err)
			assert.Error(t, handlerErr)
			assert.Equal(t, "name field is required", handlerErr.Error())
		}
	})

	t.Run("Field too short", func(t *testing.T) {
		testData := TestStruct{Name: "A", Email: "test@example.com"}
		err := validate.Struct(testData)
		if err != nil {
			handlerErr := HandleValidationErrors(err)
			assert.Error(t, handlerErr)
			assert.Equal(t, "Minimum length of the name field is 3", handlerErr.Error())
		}
	})

	t.Run("Invalid email", func(t *testing.T) {
		testData := TestStruct{Name: "Test user", Email: "invalid email"}
		err := validate.Struct(testData)
		if err != nil {
			handlerErr := HandleValidationErrors(err)
			assert.Error(t, handlerErr)
			assert.Equal(t, "email field must be a valid email", handlerErr.Error())
		}
	})

	t.Run("Valid fields", func(t *testing.T) {
		testData := TestStruct{Name: "Alice", Email: "test@example.com"}
		err := validate.Struct(testData)
		handledErr := HandleValidationErrors(err)
		assert.NoError(t, handledErr)
	})
}
