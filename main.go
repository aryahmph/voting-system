package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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
	closedAt, err := strconv.Atoi(config.Get("CLOSED_AT"))
	exception.PanicIfError(err)
	jwtConfig := configuration.JWTConfig{
		ApplicationName:    config.Get("JWT_APPLICATION_NAME"),
		SignatureKey:       []byte(config.Get("JWT_SIGNATURE_KEY")),
		ExpirationDuration: expirationConfig,
		ClosedAt:           int64(closedAt),
	}

	// Mail Config
	mailEmail := config.Get("MAIL_EMAIL")
	mailName := config.Get("MAIL_NAME")
	smtpPort, err := strconv.Atoi(config.Get("MAIL_SMTP_PORT"))
	exception.PanicIfError(err)
	mailConfig := configuration.MailConfig{
		SmtpHost: config.Get("MAIL_SMTP_HOST"),
		Name:     fmt.Sprintf("%s <%s>", mailName, mailEmail),
		Email:    mailEmail,
		Password: config.Get("MAIL_PASSWORD"),
		SmtpPort: smtpPort,
	}

	db := database.NewDatabase(dbConfig)
	validate := validator.New()

	mailService := service.NewMailServiceImpl(mailConfig)

	authService := service.NewAuthServiceImpl(jwtConfig)
	authMiddleware := middleware.NewAuthMiddleware(authService)

	voterRepository := repository.NewVoterRepositoryImpl()
	voterService := service.NewVoterServiceImpl(db, validate, voterRepository, authService)
	voterController := controller.NewVoterController(voterService, authService, mailService)

	adminRepository := repository.NewAdminRepositoryImpl()
	adminService := service.NewAdminServiceImpl(db, validate, adminRepository)
	adminController := controller.NewAdminController(adminService, voterService, authService, mailService)

	candidateRepository := repository.NewCandidateRepositoryImpl()
	candidateService := service.NewCandidateServiceImpl(db, candidateRepository)
	candidateController := controller.NewCandidateController(candidateService)

	app := fiber.New(configuration.NewFiberConfig())
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(recover.New())

	api := app.Group("/api")
	admins := api.Group("/admins", authMiddleware)
	voters := api.Group("/voters")
	candidates := api.Group("/candidates")

	api.Post("/secret-sessions", adminController.Login)

	admins.Post("/", adminController.Create)
	admins.Get("/", adminController.List)
	admins.Get("/:id", adminController.Get)
	admins.Delete("/:id", adminController.Delete)
	admins.Post("/generate-token", adminController.GenerateVoterToken)

	voters.Get("/:token", voterController.Login)
	voters.Post("/:token", voterController.Vote)

	candidates.Get("/count", candidateController.Count)

	err = app.Listen(":8080")
	exception.PanicIfError(err)
}
