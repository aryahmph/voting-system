package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"voting-system/model/domain"
)

type AdminRepository interface {
	Save(ctx context.Context, db *sqlx.DB, admin domain.Admin)
	Delete(ctx context.Context, db *sqlx.DB, id uint32)
	FindAll(ctx context.Context, db *sqlx.DB) []domain.Admin
	FindById(ctx context.Context, db *sqlx.DB, id uint32) (domain.Admin, error)
	FindByNIM(ctx context.Context, db *sqlx.DB, nim string) (domain.Admin, error)
}
