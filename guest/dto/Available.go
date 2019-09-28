package dto

import uuid "github.com/satori/go.uuid"

// Available is a message, destined to guests and sent by Hogwarts,
// indicating that there is less than 10 ongoing visits
type Available struct {
	ID               uuid.UUID `json:"id"`
	AvailableSlot    int       `json:"available_slot"`
	AvailableMessage string    `json:"available_message"`
}
