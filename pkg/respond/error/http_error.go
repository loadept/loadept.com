package respond

import "fmt"

type APIError[T string | []byte] struct {
	Message    T
	StatusCode int
}

func (e *APIError[T]) Error() string {
	return fmt.Sprintf("%v", e.Message)
}

type FieldError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (e *FieldError) Error() string {
	return e.Message
}
