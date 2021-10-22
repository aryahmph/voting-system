package configuration

import (
	"github.com/gofiber/fiber/v2"
	"voting-system/pkg/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler: exception.ErrorHandler,
	}
}