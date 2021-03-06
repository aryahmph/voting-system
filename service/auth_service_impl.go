package service

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
	"voting-system/pkg/configuration"
	"voting-system/pkg/exception"
)

type AuthServiceImpl struct {
	JWTConfig configuration.JWTConfig
}

func NewAuthServiceImpl(JWTConfig configuration.JWTConfig) *AuthServiceImpl {
	return &AuthServiceImpl{JWTConfig: JWTConfig}
}

func (service *AuthServiceImpl) GenerateToken(id uint32, role string) (string, error) {
	// Check is election has been closed or not
	now := time.Now()
	if service.JWTConfig.StartedAt >= now.Unix() && service.JWTConfig.ClosedAt <= now.Unix() {
		panic(exception.MethodNotAllowedError)
	}

	claims := configuration.AuthClaims{
		StandardClaims: jwt.StandardClaims{
			Issuer:    service.JWTConfig.ApplicationName,
			ExpiresAt: now.Add(time.Duration(service.JWTConfig.ExpirationDuration) * time.Minute).Unix(),
		},
		ID:   id,
		Role: role,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString(service.JWTConfig.SignatureKey)

	return signedString, err
}

func (service *AuthServiceImpl) ValidateToken(encodedToken string) (*jwt.Token, error) {
	parseToken, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("signing method invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, errors.New("signing method invalid")
		}
		return service.JWTConfig.SignatureKey, nil
	})

	return parseToken, err
}
