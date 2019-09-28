package dto

import "github.com/satori/go.uuid" // go get github.com/satori/go.uuid


// Slot is the message sent by Guest (via Owls)
// to the Hogwarts, announce a new visits
// from the guest
type Slot struct {
	ID       uuid.UUID `json:"id"`
	WizardID uuid.UUID `json:"wizardID"`
	SlotMessage  string    `json:"slot_message"`
}
