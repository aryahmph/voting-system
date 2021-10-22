package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"golang.org/x/crypto/bcrypt"
	"voting-system/model/domain"
	"voting-system/model/payload"
	"voting-system/pkg/exception"
	"voting-system/repository"
)

type AdminServiceImpl struct {
	DB              *sqlx.DB
	Validate        *validator.Validate
	AdminRepository repository.AdminRepository
}

func NewAdminServiceImpl(DB *sqlx.DB, validate *validator.Validate, adminRepository repository.AdminRepository) *AdminServiceImpl {
	return &AdminServiceImpl{DB: DB, Validate: validate, AdminRepository: adminRepository}
}

func (service *AdminServiceImpl) Create(ctx context.Context, request payload.CreateAdminRequest) {
	err := service.Validate.Struct(request)
	exception.PanicIfError(err)

	// Check NIM
	_, err = service.AdminRepository.FindByNIM(ctx, service.DB, request.NIM)
	if err == nil {
		err = exception.AlreadyExistError
		panic(err)
	}

	// Hashing password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	exception.PanicIfError(err)

	admin := domain.Admin{
		Name:         request.Name,
		NIM:          request.NIM,
		PasswordHash: string(passwordHash),
		Role:         "admin",
	}
	service.AdminRepository.Save(ctx, service.DB, admin)
}

func (service *AdminServiceImpl) Delete(ctx context.Context, id uint32) {
	// Check ID
	_, err := service.AdminRepository.FindById(ctx, service.DB, id)
	if err != nil {
		err = exception.NotFoundError
		panic(err)
	}

	service.AdminRepository.Delete(ctx, service.DB, id)
}

func (service *AdminServiceImpl) FindAll(ctx context.Context) []payload.GetAdminResponse {
	admins := service.AdminRepository.FindAll(ctx, service.DB)
	var responses []payload.GetAdminResponse
	for _, admin := range admins {
		responses = append(responses, payload.GetAdminResponse{
			ID:   admin.Id,
			Name: admin.Name,
			NIM:  admin.NIM,
			Role: admin.Role,
		})
	}
	return responses
}

func (service *AdminServiceImpl) FindById(ctx context.Context, id uint32) payload.GetAdminResponse {
	admin, err := service.AdminRepository.FindById(ctx, service.DB, id)
	if err != nil {
		err = exception.NotFoundError
		panic(err)
	}

	return payload.GetAdminResponse{
		ID:           id,
		Name:         admin.Name,
		NIM:          admin.NIM,
		PasswordHash: admin.PasswordHash,
		Role:         admin.Role,
	}
}
