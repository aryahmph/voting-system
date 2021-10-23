package service

import (
	"context"
	"voting-system/model/payload"
)

type VoterService interface {
	GenerateVote(ctx context.Context, request payload.GenerateVoteRequest) payload.GenerateVoteResponse
	Vote(ctx context.Context, request payload.VoteRequest)
}