package dto

import uuid "github.com/satori/go.uuid"

// Available is a message, destined to guests,
// indicating that there is less than 10 ongoing visits
type Available struct {
	ID               uuid.UUID `json:"id"`
	AvailableSlot    int       `json:"availableSlot"`
	AvailableMessage string    `json:"Available_message"`
}

