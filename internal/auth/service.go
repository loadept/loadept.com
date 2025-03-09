package auth

import (
	"github.com/loadept/loadept.com/internal/auth/jwt"
	"github.com/loadept/loadept.com/internal/config"
)

type TokenService interface {
	CreateToken(string, bool) (string, error)
	ValidateToken(string) (*jwt.TokenClaims, error)
}

func NewAuthService(service string) TokenService {
	switch service {
	case "JWT":
		return &jwt.JWT{Secret: config.Env.SECRET_KEY}
	default:
		return nil
	}
}
