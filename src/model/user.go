package model

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID `db:"id"`
	FirstName string    `db=:"first_name" json:"first_name"`
	LastName  string    `db=:"last_name" json:"last_name"`
}
