package service

import "github.com/golang-jwt/jwt"

type AuthService interface {
	GenerateToken(id uint32, role string) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
}
