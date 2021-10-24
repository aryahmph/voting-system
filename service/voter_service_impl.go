package service

import (
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/jmoiron/sqlx"
	"voting-system/model/payload"
	"voting-system/pkg/exception"
	"voting-system/repository"
)

type VoterServiceImpl struct {
	DB              *sqlx.DB
	Validate        *validator.Validate
	VoterRepository repository.VoterRepository
	AuthService     AuthService
}

func NewVoterServiceImpl(DB *sqlx.DB, validate *validator.Validate, voterRepository repository.VoterRepository, authService AuthService) *VoterServiceImpl {
	return &VoterServiceImpl{DB: DB, Validate: validate, VoterRepository: voterRepository, AuthService: authService}
}

func (service *VoterServiceImpl) GenerateVote(ctx context.Context, request payload.GenerateVoteRequest) payload.GenerateVoteResponse {
	err := service.Validate.Struct(request)
	exception.PanicIfError(err)

	// Check NIM
	voter, err := service.VoterRepository.FindByNIM(ctx, service.DB, request.NIM)
	if err != nil {
		panic(exception.NotFoundError)
	}
	// Check has voted or not
	if voter.AdminID.Valid {
		panic(exception.AlreadyExistError)
	}

	// Generate token
	token, err := service.AuthService.GenerateToken(voter.ID, "voter")
	exception.PanicIfError(err)

	// Mark admin who generate
	service.VoterRepository.Update(ctx, service.DB, voter.ID, request.AdminID)

	return payload.GenerateVoteResponse{
		Name:  voter.Name,
		Email: voter.Email,
		Token: token,
	}
}

func (service *VoterServiceImpl) Vote(ctx context.Context, request payload.VoteRequest) payload.VoteResponse {
	err := service.Validate.Struct(request)
	exception.PanicIfError(err)

	voter, err := service.VoterRepository.FindByID(ctx, service.DB, request.ID)
	if err != nil {
		panic(exception.NotFoundError)
	}

	if voter.CandidateID.Valid {
		panic(exception.AlreadyExistError)
	}

	service.VoterRepository.UpdateCandidate(ctx, service.DB, request.ID, request.CandidateID)
	return payload.VoteResponse{
		Name:  voter.Name,
		NIM:   voter.NIM,
		Email: voter.Email,
	}
}

func (service *VoterServiceImpl) FindByID(ctx context.Context, id uint32) payload.VoterLoginResponse {
	voter, err := service.VoterRepository.FindByID(ctx, service.DB, id)
	if err != nil {
		panic(exception.NotFoundError)
	}

	hasVoted := false
	if voter.CandidateID.Valid {
		hasVoted = true
	}

	return payload.VoterLoginResponse{
		ID:       voter.ID,
		Name:     voter.Name,
		NIM:      voter.NIM,
		Email:    voter.Email,
		HasVoted: hasVoted,
	}
}
