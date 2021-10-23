package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"net/http"
	"voting-system/model/payload"
	"voting-system/pkg/exception"
	"voting-system/service"
)

type VoterController struct {
	VoterService service.VoterService
	AuthService  service.AuthService
}

func NewVoterController(voterService service.VoterService, authService service.AuthService) *VoterController {
	return &VoterController{VoterService: voterService, AuthService: authService}
}

func (controller *VoterController) Vote(ctx *fiber.Ctx) error {
	token := ctx.Params("token")

	// Parse request
	var request payload.VoteRequest
	err := ctx.BodyParser(&request)
	exception.PanicIfError(err)

	// Validate Token
	validateToken, err := controller.AuthService.ValidateToken(token)
	if err != nil {
		panic(exception.UnauthorizedError)
	}
	claims, ok := validateToken.Claims.(jwt.MapClaims)
	if !ok || !validateToken.Valid {
		panic(exception.UnauthorizedError)
	}

	request.ID = uint32(claims["id"].(float64))
	role := claims["role"].(string)
	if role != "voter" {
		panic(exception.UnauthorizedError)
	}

	controller.VoterService.Vote(ctx.Context(), request)

	return ctx.JSON(payload.WebResponse{
		Code:   200,
		Status: http.StatusText(200),
		Data:   nil,
	})
}