package middleware

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"strings"
	"voting-system/model/payload"
	"voting-system/pkg/exception"
	"voting-system/service"
)

func NewAuthMiddleware(service service.AuthService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			panic(exception.UnauthorizedError)
		}

		tokenSlice := strings.Split(authHeader, " ")
		if len(tokenSlice) != 2 {
			panic(exception.UnauthorizedError)
		}

		validateToken, err := service.ValidateToken(tokenSlice[1])
		if err != nil {
			panic(exception.UnauthorizedError)
		}

		claims, ok := validateToken.Claims.(jwt.MapClaims)
		if !ok || !validateToken.Valid {
			panic(exception.UnauthorizedError)
		}

		id := uint32(claims["id"].(float64))
		role := claims["role"].(string)
		authMiddleware := payload.AuthMiddleware{
			ID:   id,
			Role: role,
		}

		ctxCustom := context.WithValue(ctx.Context(), "userAuth", authMiddleware)
		ctx.SetUserContext(ctxCustom)
		return ctx.Next()
	}
}
