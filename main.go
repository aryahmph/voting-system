package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"strconv"
	"voting-system/controller"
	"voting-system/pkg/configuration"
	"voting-system/pkg/database"
	"voting-system/pkg/exception"
	"voting-system/repository"
	"voting-system/service"
)

func main() {
	config := configuration.NewConfigurationImpl(".env")

	poolMinConfig, err := strconv.Atoi(config.Get("MYSQL_POOL_MIN"))
	exception.PanicIfError(err)
	poolMaxConfig, err := strconv.Atoi(config.Get("MYSQL_POOL_MAX"))
	exception.PanicIfError(err)
	dbConfig := configuration.DbConfig{
		User:     config.Get("MYSQL_USER"),
		Password: config.Get("MYSQL_PASSWORD"),
		Host:     config.Get("MYSQL_HOST"),
		Port:     config.Get("MYSQL_PORT"),
		Database: config.Get("MYSQL_DATABASE"),
		PoolMin:  poolMinConfig,
		PoolMax:  poolMaxConfig,
	}

	db := database.NewDatabase(dbConfig)
	validate := validator.New()

	adminRepository := repository.NewAdminRepositoryImpl()
	adminService := service.NewAdminServiceImpl(db, validate, adminRepository)
	adminController := controller.NewAdminController(adminService)

	app := fiber.New(configuration.NewFiberConfig())
	app.Use(logger.New())
	app.Use(recover.New())

	adminController.Route(app)

	err = app.Listen(":8080")
	exception.PanicIfError(err)
}
