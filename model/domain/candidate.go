package domain

type Candidate struct {
	ID         uint32
	Name       string
	VotesCount uint32 `db:"votes_count"`
}
