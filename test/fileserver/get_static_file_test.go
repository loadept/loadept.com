package test

import (
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/loadept/loadept.com/api"
)

func TestServeStaticFile(t *testing.T) {
	tmpDir := t.TempDir()
	err := os.WriteFile(filepath.Join(tmpDir, "testfile.txt"), []byte("Hello, test file!\n"), 0644)
	if err != nil {
		t.Fatalf("Error to create temp file: %v", err)
	}
	err = os.WriteFile(filepath.Join(tmpDir, ".hidden"), []byte("Hello, hidden test file!\n"), 0644)
	if err != nil {
		t.Fatalf("Error to create temp file: %v", err)
	}

	t.Run("Valid File", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/testfile.txt", nil)
		rr := httptest.NewRecorder()

		handler := api.ServeStaticFile(filepath.Join(tmpDir, "testfile.txt"))
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", rr.Code)
		}

		expectedBody := "Hello, test file!\n"
		if rr.Body.String() != expectedBody {
			t.Errorf("Expected body %q, got %q", expectedBody, rr.Body.String())
		}
	})

	t.Run("Invalid File", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/nonexistent.txt", nil)
		rr := httptest.NewRecorder()

		handler := api.ServeStaticFile(filepath.Join(tmpDir, "testfile.txt"))
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("Expected status 404, got %d", rr.Code)
		}

		expectedBody := "404 page not found\n"
		if rr.Body.String() != expectedBody {
			t.Errorf("Expected body %q, got %q", expectedBody, rr.Body.String())
		}
	})

	t.Run("Invalid method", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/testfile.txt", nil)
		rr := httptest.NewRecorder()

		handler := api.ServeStaticFile(filepath.Join(tmpDir, "testfile.txt"))
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusMethodNotAllowed {
			t.Errorf("Expected status 405, got %d", rr.Code)
		}

		expectedBody := "This method is not allowed\n"
		if rr.Body.String() != expectedBody {
			t.Errorf("Expected body %q, got %q", expectedBody, rr.Body.String())
		}
	})

	t.Run("HEAD method", func(t *testing.T) {
		req := httptest.NewRequest("HEAD", "/testfile.txt", nil)
		rr := httptest.NewRecorder()

		handler := api.ServeStaticFile(filepath.Join(tmpDir, "testfile.txt"))
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusOK {
			t.Errorf("Expected status 200, got %d", rr.Code)
		}
	})

	t.Run("Directory traversal", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/../../../../../../etc/passwd", nil)
		rr := httptest.NewRecorder()

		handler := api.ServeStaticFile(filepath.Join(tmpDir, "testfile.txt"))
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("Expected status 404, got %d", rr.Code)
		}
	})

	t.Run("Hidden file", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/.hidden", nil)
		rr := httptest.NewRecorder()

		handler := api.ServeStaticFile(filepath.Join(tmpDir, "testfile.txt"))
		handler.ServeHTTP(rr, req)

		if rr.Code != http.StatusNotFound {
			t.Errorf("Expected status 404, got %d", rr.Code)
		}
	})
}
