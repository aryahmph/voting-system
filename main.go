package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"strconv"
	"voting-system/controller"
	"voting-system/middleware"
	"voting-system/pkg/configuration"
	"voting-system/pkg/database"
	"voting-system/pkg/exception"
	"voting-system/repository"
	"voting-system/service"
)

func main() {
	config := configuration.NewConfigurationImpl(".env")

	// Db config
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

	// JWT Config
	expirationConfig, err := strconv.Atoi(config.Get("JWT_EXPIRATION_DURATION"))
	exception.PanicIfError(err)
	jwtConfig := configuration.JWTConfig{
		ApplicationName:    config.Get("JWT_APPLICATION_NAME"),
		SignatureKey:       config.Get("JWT_SIGNATURE_KEY"),
		ExpirationDuration: expirationConfig,
	}

	db := database.NewDatabase(dbConfig)
	validate := validator.New()

	authService := service.NewAuthServiceImpl(jwtConfig)
	authMiddleware := middleware.NewAuthMiddleware(authService)

	adminRepository := repository.NewAdminRepositoryImpl()
	adminService := service.NewAdminServiceImpl(db, validate, adminRepository)
	adminController := controller.NewAdminController(adminService, authService)

	app := fiber.New(configuration.NewFiberConfig())
	app.Use(logger.New())
	app.Use(recover.New())

	api := app.Group("/api")
	admins := api.Group("/admins", authMiddleware)

	admins.Post("/", adminController.Create)
	admins.Get("/", adminController.List)
	admins.Delete("/:id", adminController.Delete)

	err = app.Listen(":8080")
	exception.PanicIfError(err)
}
