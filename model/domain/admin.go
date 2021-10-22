package domain

import "time"

type Admin struct {
	Id           uint32
	Name         string
	NIM          string
	PasswordHash string `db:"password_hash"`
	Role         string
	DeletedAt    *time.Time `db:"deleted_at"`
}
