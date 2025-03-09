package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/loadept/loadept.com/internal/auth/jwt"
	"github.com/loadept/loadept.com/pkg/respond"
)

const AuthContextKey = "AuthContext"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerToken := r.Header.Get("Authorization")
		if !strings.HasPrefix(headerToken, "Bearer ") {
			respond.JSON(w, respond.Map{
				"detail": "Resource not found",
			}, http.StatusNotFound)
			return
		}

		token := strings.TrimPrefix(headerToken, "Bearer ")

		claims, err := jwt.JWTAuth().ValidateToken(token)
		if err != nil {
			respond.JSON(w, respond.Map{
				"detail": "The token is not valid",
			}, http.StatusUnauthorized)
			return
		}

		if !claims.IsAdmin {
			respond.JSON(w, respond.Map{
				"detail": "The token is not valid",
			}, http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), AuthContextKey, respond.Map{
			"user_id":  claims.UserID,
			"is_admin": claims.IsAdmin,
		})

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
