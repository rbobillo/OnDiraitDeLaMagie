package dao

import uuid "github.com/satori/go.uuid"

// Action is the content for the Hogwarts Inventory DB
type Action struct {
	ID         uuid.UUID `json:"id"`
	Wizard_id  uuid.UUID `json:"wizard_id"`
	Action     string    `json:"action"`
	Status     string    `json:"status"`
}