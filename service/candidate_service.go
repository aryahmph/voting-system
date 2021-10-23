package service

import (
	"context"
	"voting-system/model/payload"
)

type CandidateService interface {
	CountVotes(ctx context.Context) []payload.CountVotesResponse
}
