package dto

import "github.com/satori/go.uuid" // go get github.com/satori/go.uuid

// TODO: by who is sent the eligible mail ?
// Eligible is the message sent by magic? (via Owls)
// to families, announce a
// Wizards is ready to go to Hogwarts
type Eligible struct {
	ID       uuid.UUID `json:"id"`
	WizardID uuid.UUID `json:"wizardID"`
	Message  string    `json:"message"`
}