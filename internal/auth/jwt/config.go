package jwt

import "github.com/golang-jwt/jwt"

type TokenClaims struct {
	UserID  string `json:"user_id"`
	IsAdmin bool   `json:"is_admin"`
	jwt.StandardClaims
}
