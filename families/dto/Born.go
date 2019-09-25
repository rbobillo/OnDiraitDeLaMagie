package dto

import "github.com/satori/go.uuid" // go get github.com/satori/go.uuid


// Birth is the message sent by Families (via Owls)
// to the Ministry, announce a new born
// Wizards in the families
type Birth struct {
	ID          uuid.UUID `json:"id"`
	WizardID    uuid.UUID `json:"wizard_id"`
	BornMessage string    `json:"born_message"`
}