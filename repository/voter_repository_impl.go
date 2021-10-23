package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"voting-system/model/domain"
)

type VoterRepositoryImpl struct {
}

func NewVoterRepositoryImpl() *VoterRepositoryImpl {
	return &VoterRepositoryImpl{}
}

func (repository *VoterRepositoryImpl) FindByID(ctx context.Context, db *sqlx.DB, id uint32) (domain.Voter, error) {
	voter := domain.Voter{}
	err := db.GetContext(ctx, &voter,
		"SELECT id, candidate_id, admin_id, name, nim, email\nFROM voters\nWHERE id = ?\nLIMIT 1;", id)

	if err != nil {
		return domain.Voter{}, err
	}
	return voter, nil
}

func (repository *VoterRepositoryImpl) FindByNIM(ctx context.Context, db *sqlx.DB, nim string) (domain.Voter, error) {
	voter := domain.Voter{}
	err := db.GetContext(ctx, &voter,
		"SELECT id, candidate_id, admin_id, name, nim, email\nFROM voters\nWHERE nim = ?\nLIMIT 1;", nim)

	if err != nil {
		return domain.Voter{}, err
	}
	return voter, nil
}

func (repository *VoterRepositoryImpl) Update(ctx context.Context, db *sqlx.DB, id uint32, adminID uint32) {
	db.MustExecContext(ctx,
		"UPDATE voters\nSET admin_id=?,\n    updated_at=CURRENT_TIMESTAMP\nWHERE id = ?",
		adminID, id)
}

func (repository *VoterRepositoryImpl) UpdateCandidate(ctx context.Context, db *sqlx.DB, id uint32, candidateID uint32) {
	db.MustExecContext(ctx,
		"UPDATE voters\nSET candidate_id=?,\n    updated_at=CURRENT_TIMESTAMP\nWHERE id = ?",
		candidateID, id)
}
