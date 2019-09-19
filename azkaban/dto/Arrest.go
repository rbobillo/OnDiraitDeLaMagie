package dto

import uuid "github.com/satori/go.uuid"

// Arrest is the message from Ministry
// ordering Azkaban to put a wizard in jail
type Arrest struct {
	ID 		 uuid.UUID `json:"id"`
	WizardID uuid.UUID `json:"wizard_id"`
}