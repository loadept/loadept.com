package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/loadept/loadept.com/internal/config"
	"github.com/loadept/loadept.com/pkg/logger"
)

func TestLoggerMiddleware(t *testing.T) {
	err := os.MkdirAll("logs", os.ModePerm)
	if err != nil {
		t.Fatalf("No se pudo crear el directorio logs: %v", err)
	}

	os.Setenv("DEBUG", "false")
	defer os.Unsetenv("DEBUG")
	config.LoadEnviron()

	defer os.RemoveAll("logs")
	logger.NewLogger()

	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	handler := LoggerMiddleware(testHandler)

	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("Expected status 200, got %d", rr.Code)
	}

	expectedBody := "OK"
	if rr.Body.String() != expectedBody {
		t.Errorf("Expected body %s, got %q", expectedBody, rr.Body.String())
	}
}
