package dto

import uuid "github.com/satori/go.uuid"

// Available is a message, destined to guests,
// indicating that there is less than 10 ongoing visits
type Available struct {
	ID               uuid.UUID `json:"id"`
	GuestID 		 uuid.UUID `json:"guest_id"`
	AvailableMessage string    `json:"available_message"`
}

