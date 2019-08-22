package dao

import uuid "github.com/satori/go.uuid"

// Student is the content of table students
// from hogwartsinventory DB
type Student struct {
	ID      uuid.UUID `json:"id"`
	MagicID uuid.UUID `json:"magic_id"`
	House   string    `json:"house"`
	Status  string    `json:"status"`
}