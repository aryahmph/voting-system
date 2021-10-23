package exception

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"voting-system/model/payload"
)

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	// Unauthorized error
	if ok := errors.Is(UnauthorizedError, err); ok {
		return ctx.Status(http.StatusUnauthorized).JSON(payload.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: http.StatusText(http.StatusUnauthorized),
			Error:  err.Error(),
		})
	}

	// NotFound error
	if ok := errors.Is(NotFoundError, err); ok {
		return ctx.Status(http.StatusNotFound).JSON(payload.WebResponse{
			Code:   http.StatusNotFound,
			Status: http.StatusText(http.StatusNotFound),
			Error:  err.Error(),
		})
	}

	// AlreadyExist error
	if ok := errors.Is(AlreadyExistError, err); ok {
		return ctx.Status(http.StatusConflict).JSON(payload.WebResponse{
			Code:   http.StatusConflict,
			Status: http.StatusText(http.StatusConflict),
			Error:  err.Error(),
		})
	}

	// MethodNotAllowed error
	if ok := errors.Is(MethodNotAllowedError, err); ok {
		return ctx.Status(http.StatusMethodNotAllowed).JSON(payload.WebResponse{
			Code:   http.StatusMethodNotAllowed,
			Status: http.StatusText(http.StatusMethodNotAllowed),
			Error:  err.Error(),
		})
	}

	// Validation error
	if exception, ok := err.(validator.ValidationErrors); ok {
		return ctx.Status(http.StatusUnprocessableEntity).JSON(payload.WebResponse{
			Code:   http.StatusUnprocessableEntity,
			Status: http.StatusText(http.StatusUnprocessableEntity),
			Error:  exception.Error(),
		})
	}

	// Internal Error
	return ctx.Status(500).JSON(payload.WebResponse{
		Code:   500,
		Status: http.StatusText(500),
		Error:  err.Error(),
	})
}
