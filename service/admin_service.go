package service

import (
	"context"
	"voting-system/model/payload"
)

type AdminService interface {
	Create(ctx context.Context, request payload.CreateAdminRequest)
	Delete(ctx context.Context, id uint32)
	FindAll(ctx context.Context) []payload.GetAdminResponse
	FindById(ctx context.Context, id uint32) payload.GetAdminResponse
}
