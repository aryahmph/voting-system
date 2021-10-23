package domain

type Candidate struct {
	ID         uint32
	VotesCount uint32 `db:"votes_count"`
}
