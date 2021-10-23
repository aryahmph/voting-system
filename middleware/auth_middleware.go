package middleware

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"net/http"
	"strings"
	"voting-system/model/payload"
	"voting-system/service"
)

func NewAuthMiddleware(service service.AuthService) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		if !strings.Contains(authHeader, "Bearer") {
			fmt.Println("error contains")
			return ctx.Status(http.StatusUnauthorized).JSON(payload.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
				Error:  "token is not valid",
			})
		}

		tokenSlice := strings.Split(authHeader, " ")
		if len(tokenSlice) != 2 {
			fmt.Println("error slice")
			return ctx.Status(http.StatusUnauthorized).JSON(payload.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
				Error:  "token is not valid",
			})
		}

		validateToken, err := service.ValidateToken(tokenSlice[1])
		if err != nil {
			fmt.Println("error hehe", err)
			return ctx.Status(http.StatusUnauthorized).JSON(payload.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
				Error:  "token is not valid",
			})
		}

		claims, ok := validateToken.Claims.(jwt.MapClaims)
		if !ok || !validateToken.Valid {
			return ctx.Status(http.StatusUnauthorized).JSON(payload.WebResponse{
				Code:   http.StatusUnauthorized,
				Status: http.StatusText(http.StatusUnauthorized),
				Error:  "token is not valid",
			})
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
