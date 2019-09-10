package dto

import uuid "github.com/satori/go.uuid"

// Alert is the action of
// the Villains attacking Hogwarts
// quick/strong determine how efficient
// the attack will be
type Slot struct {
	ID       uuid.UUID `json:"id"`
	WizardID uuid.UUID `json:"slot_id"`
}

