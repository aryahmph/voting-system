package configuration

import "github.com/golang-jwt/jwt"

type JWTConfig struct {
	ApplicationName    string
	SignatureKey       string
	ExpirationDuration int
}

type AuthClaims struct {
	jwt.StandardClaims
	ID   uint32 `json:"id"`
	Role string `json:"role"`
}
