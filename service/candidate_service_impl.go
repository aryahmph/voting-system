package service

import (
	"context"
	"github.com/jmoiron/sqlx"
	"voting-system/model/payload"
	"voting-system/repository"
)

type CandidateServiceImpl struct {
	DB                  *sqlx.DB
	CandidateRepository repository.CandidateRepository
}

func NewCandidateServiceImpl(DB *sqlx.DB, candidateRepository repository.CandidateRepository) *CandidateServiceImpl {
	return &CandidateServiceImpl{DB: DB, CandidateRepository: candidateRepository}
}

func (service *CandidateServiceImpl) CountVotes(ctx context.Context) []payload.CountVotesResponse {
	candidates := service.CandidateRepository.Count(ctx, service.DB)
	var responses []payload.CountVotesResponse
	for _, candidate := range candidates {
		responses = append(responses, payload.CountVotesResponse{
			ID:    candidate.ID,
			Name:  candidate.Name,
			Votes: candidate.VotesCount,
		})
	}
	return responses
}
