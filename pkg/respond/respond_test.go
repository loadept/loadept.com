package respond

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJSON(t *testing.T) {
	t.Run("Sucessful response", func(t *testing.T) {
		w := httptest.NewRecorder()

		response := Map{"message": "success"}
		statusCode := http.StatusOK

		JSON(w, response, statusCode)

		assert.Equal(t, statusCode, w.Code, "Status code should be 200")
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"), "Content-Type should be application/json")

		expectedBody := `{"message":"success"}` + "\n"
		assert.Equal(t, expectedBody, w.Body.String(), "The response body does not match")
	})

	t.Run("JSON encoding error", func(t *testing.T) {
		w := httptest.NewRecorder()

		invalidResponse := make(chan int)
		statusCode := http.StatusOK

		JSON(w, invalidResponse, statusCode)

		assert.Equal(t, http.StatusInternalServerError, w.Code, "Status code should be 500")
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"), "Content-Type should be application/json")

		expectedBody := `{"detail":"Internal server error"}` + "\n"
		assert.Equal(t, expectedBody, w.Body.String(), "The response body does not match")
	})
}
