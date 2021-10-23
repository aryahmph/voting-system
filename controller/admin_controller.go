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
	AuthService  service.AuthService
}

func NewAdminController(adminService service.AdminService, authService service.AuthService) *AdminController {
	return &AdminController{AdminService: adminService, AuthService: authService}
}

func (controller *AdminController) Create(ctx *fiber.Ctx) error {
	userAuth := ctx.UserContext().Value("userAuth").(payload.AuthMiddleware)
	if userAuth.Role != "super-admin" {
		panic(exception.UnauthorizedError)
	}
	_ = controller.AdminService.FindById(ctx.Context(), userAuth.ID)

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
	userAuth := ctx.UserContext().Value("userAuth").(payload.AuthMiddleware)
	if userAuth.Role != "super-admin" {
		panic(exception.UnauthorizedError)
	}
	_ = controller.AdminService.FindById(ctx.Context(), userAuth.ID)

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
	userAuth := ctx.UserContext().Value("userAuth").(payload.AuthMiddleware)
	if userAuth.Role != "super-admin" {
		panic(exception.UnauthorizedError)
	}
	_ = controller.AdminService.FindById(ctx.Context(), userAuth.ID)

	responses := controller.AdminService.FindAll(ctx.Context())
	return ctx.JSON(payload.WebResponse{
		Code:   200,
		Status: http.StatusText(200),
		Data:   responses,
	})
}
