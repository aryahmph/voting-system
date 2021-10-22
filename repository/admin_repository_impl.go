package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"voting-system/model/domain"
	"voting-system/pkg/exception"
)

type AdminRepositoryImpl struct {
}

func NewAdminRepositoryImpl() *AdminRepositoryImpl {
	return &AdminRepositoryImpl{}
}

func (repository *AdminRepositoryImpl) Save(ctx context.Context, db *sqlx.DB, admin domain.Admin) {
	db.MustExecContext(ctx,
		"INSERT INTO admins(name, nim, password_hash, role) VALUE (?, ?, ?, ?)",
		admin.Name, admin.NIM, admin.PasswordHash, admin.Role)
}

func (repository *AdminRepositoryImpl) Delete(ctx context.Context, db *sqlx.DB, id uint32) {
	db.MustExecContext(ctx, "UPDATE admins\nSET deleted_at = CURRENT_TIMESTAMP\nWHERE id = ?", id)
}

func (repository *AdminRepositoryImpl) FindAll(ctx context.Context, db *sqlx.DB) []domain.Admin {
	var admins []domain.Admin
	err := db.SelectContext(ctx, &admins,
		"SELECT id, name, nim, password_hash, role, deleted_at\nFROM admins\nWHERE deleted_at IS NULL")

	exception.PanicIfError(err)
	return admins
}

func (repository *AdminRepositoryImpl) FindById(ctx context.Context, db *sqlx.DB, id uint32) (domain.Admin, error) {
	admin := domain.Admin{}
	err := db.GetContext(ctx, &admin,
		"SELECT id, name, nim, password_hash, role, deleted_at\nFROM admins\nWHERE id = ?\n  AND deleted_at IS NULL\nLIMIT 1", id)

	if err != nil {
		return domain.Admin{}, err
	}
	return admin, nil
}

func (repository *AdminRepositoryImpl) FindByNIM(ctx context.Context, db *sqlx.DB, nim string) (domain.Admin, error) {
	admin := domain.Admin{}
	err := db.GetContext(ctx, &admin,
		"SELECT id, name, nim, password_hash, role, deleted_at\nFROM admins\nWHERE nim = ?\n  AND deleted_at IS NULL\nLIMIT 1", nim)

	if err != nil {
		return domain.Admin{}, err
	}
	return admin, nil
}
