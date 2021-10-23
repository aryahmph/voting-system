package controller

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"voting-system/model/payload"
	"voting-system/service"
)

type CandidateController struct {
	CandidateService service.CandidateService
}

func NewCandidateController(candidateService service.CandidateService) *CandidateController {
	return &CandidateController{CandidateService: candidateService}
}

func (controller *CandidateController) Count(ctx *fiber.Ctx) error {
	responses := controller.CandidateService.CountVotes(ctx.Context())
	return ctx.JSON(payload.WebResponse{
		Code:   200,
		Status: http.StatusText(200),
		Data:   responses,
	})
}
