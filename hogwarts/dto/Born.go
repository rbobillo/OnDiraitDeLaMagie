package dto

import "github.com/satori/go.uuid" // go get github.com/satori/go.uuid


// Birth is the message sent by Ministry (via Owls)
// to the Hogwarts, announce a new born
// Wizards is ready to be enrolled
type Birth struct {
	ID          uuid.UUID `json:"id"`
	WizardID    uuid.UUID `json:"wizard_id"`
	BornMessage string    `json:"born_message"`
}