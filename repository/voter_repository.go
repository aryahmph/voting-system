package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"voting-system/model/domain"
)

type VoterRepository interface {
	FindByID(ctx context.Context, db *sqlx.DB, id uint32) (domain.Voter, error)
	FindByNIM(ctx context.Context, db *sqlx.DB, nim string) (domain.Voter, error)
	Update(ctx context.Context, db *sqlx.DB, id uint32, adminID uint32)
	UpdateCandidate(ctx context.Context, db *sqlx.DB, id uint32, candidateID uint32)
}
