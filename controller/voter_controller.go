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
	MailService  service.MailService
}

func NewVoterController(voterService service.VoterService, authService service.AuthService, mailService service.MailService) *VoterController {
	return &VoterController{VoterService: voterService, AuthService: authService, MailService: mailService}
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

func (controller *VoterController) Login(ctx *fiber.Ctx) error {
	token := ctx.Params("token")

	// Validate Token
	validateToken, err := controller.AuthService.ValidateToken(token)
	if err != nil {
		panic(exception.UnauthorizedError)
	}
	claims, ok := validateToken.Claims.(jwt.MapClaims)
	if !ok || !validateToken.Valid {
		panic(exception.UnauthorizedError)
	}

	id := uint32(claims["id"].(float64))
	role := claims["role"].(string)
	if role != "voter" {
		panic(exception.UnauthorizedError)
	}

	response := controller.VoterService.FindByID(ctx.Context(), id)

	return ctx.JSON(payload.WebResponse{
		Code:   200,
		Status: http.StatusText(200),
		Data:   response,
	})
}
