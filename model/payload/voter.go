package payload

type GetVoterResponse struct {
	ID          uint32 `json:"id"`
	CandidateID uint32 `json:"candidate_id"`
	AdminID     uint32 `json:"admin_id"`
	Name        string `json:"name"`
	NIM         string `json:"nim"`
	Email       string `json:"email"`
}

type GenerateVoteRequest struct {
	AdminID uint32 `validate:"required"`
	NIM     string `json:"nim" validate:"required"`
}

type GenerateVoteResponse struct {
	Token string `json:"token"`
}

type VoteRequest struct {
	ID          uint32 `validate:"required"`
	CandidateID uint32 `json:"candidate_id" validate:"required"`
}

type VoterLoginResponse struct {
	ID       uint32 `json:"id"`
	Name     string `json:"name"`
	NIM      string `json:"nim"`
	Email    string `json:"email"`
	HasVoted bool   `json:"has_voted"`
}
