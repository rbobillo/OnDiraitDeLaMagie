package dto

import uuid "github.com/satori/go.uuid"

// Available is a message, destined to guests,
// indicating that there is less than 10 ongoing visits
type Arrested struct {
	ID              uuid.UUID `json:"id"`
	WizardID        uuid.UUID `json:"wizard_id"`
	ArrestedMessage string    `json:"arrested_message"`
}

