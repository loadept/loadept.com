package middleware

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/loadept/loadept.com/internal/auth/jwt"
	"github.com/stretchr/testify/assert"
)

type mockJWT struct{}

func (m *mockJWT) ValidateToken(token string) (*jwt.TokenClaims, error) {
	if token == "valid-token" {
		return &jwt.TokenClaims{
			UserID:  "123",
			IsAdmin: true,
		}, nil
	}
	return nil, fmt.Errorf("invalid token")
}

func (m *mockJWT) CreateToken(string, bool) (string, error) { return "", nil }

func TestAuthMiddleware(t *testing.T) {
	tests := []struct {
		name           string
		token          string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "Valid token",
			token:          "Bearer valid-token",
			expectedStatus: http.StatusOK,
			expectedBody:   "OK",
		},
		{
			name:           "TokenNotProvided",
			token:          "",
			expectedStatus: http.StatusNotFound,
			expectedBody:   `{"detail":"Resource not found"}` + "\n",
		},
		{
			name:           "Valid token",
			token:          "Bearer valid-token",
			expectedStatus: http.StatusOK,
			expectedBody:   "OK",
		},
		{
			name:           "Invalid token",
			token:          "Bearer invalid-token",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"detail":"The token is not valid"}` + "\n",
		},
		{
			name:           "Valid token but is not admin",
			token:          "Bearer valid-token-not-admin",
			expectedStatus: http.StatusUnauthorized,
			expectedBody:   `{"detail":"The token is not valid"}` + "\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", "/protected", nil)
			assert.NoError(t, err)

			if tt.token != "" {
				req.Header.Set("Authorization", tt.token)
			}

			rr := httptest.NewRecorder()

			handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("OK"))
			})

			authMiddleware := NewAuthMiddleware(&mockJWT{})
			authMiddleware(handler).ServeHTTP(rr, req)

			assert.Equal(t, tt.expectedStatus, rr.Code)
			assert.Equal(t, tt.expectedBody, rr.Body.String())
		})
	}
}
