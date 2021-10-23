package domain

import "database/sql"

type Voter struct {
	ID               uint32
	CandidateID      sql.NullInt32 `db:"candidate_id"`
	AdminID          sql.NullInt32 `db:"admin_id"`
	Name, NIM, Email string
}
