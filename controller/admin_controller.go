package controller

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
	"voting-system/model/payload"
	"voting-system/pkg/exception"
	"voting-system/service"
)

type AdminController struct {
	AdminService service.AdminService
}

func NewAdminController(adminService service.AdminService) *AdminController {
	return &AdminController{AdminService: adminService}
}

func (controller *AdminController) Route(app *fiber.App) {
	app.Post("/api/admins", controller.Create)
	app.Get("/api/admins", controller.List)
	app.Delete("/api/admins/:id", controller.Delete)
}

func (controller *AdminController) Create(ctx *fiber.Ctx) error {
	var request payload.CreateAdminRequest
	err := ctx.BodyParser(&request)
	exception.PanicIfError(err)

	controller.AdminService.Create(ctx.Context(), request)
	return ctx.JSON(payload.WebResponse{
		Code:   200,
		Status: http.StatusText(200),
		Data:   nil,
	})
}

func (controller *AdminController) Delete(ctx *fiber.Ctx) error {
	id, err := ctx.ParamsInt("id")
	exception.PanicIfError(err)

	controller.AdminService.Delete(ctx.Context(), uint32(id))
	return ctx.JSON(payload.WebResponse{
		Code:   200,
		Status: http.StatusText(200),
		Data:   nil,
	})
}

func (controller *AdminController) List(ctx *fiber.Ctx) error {
	responses := controller.AdminService.FindAll(ctx.Context())
	return ctx.JSON(payload.WebResponse{
		Code:   200,
		Status: http.StatusText(200),
		Data:   responses,
	})
}
