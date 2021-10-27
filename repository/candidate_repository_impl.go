package repository

import (
	"context"
	"github.com/jmoiron/sqlx"
	"voting-system/model/domain"
	"voting-system/pkg/exception"
)

type CandidateRepositoryImpl struct {
}

func NewCandidateRepositoryImpl() *CandidateRepositoryImpl {
	return &CandidateRepositoryImpl{}
}

func (repository *CandidateRepositoryImpl) Count(ctx context.Context, db *sqlx.DB) []domain.Candidate {
	var candidates []domain.Candidate
	err := db.SelectContext(ctx, &candidates,
		"SELECT candidates.id, candidates.name, COUNT(voters.id) as votes_count\nFROM candidates\nLEFT JOIN voters ON candidates.id = voters.candidate_id\nGROUP BY candidates.id")

	exception.PanicIfError(err)
	return candidates
}
