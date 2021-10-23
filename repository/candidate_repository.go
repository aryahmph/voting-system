package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"voting-system/model/domain"
)

type CandidateRepository interface {
	Count(ctx context.Context, db *sqlx.DB) []domain.Candidate
}
