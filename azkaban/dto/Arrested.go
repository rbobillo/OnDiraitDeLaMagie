package dto

import uuid "github.com/satori/go.uuid"

// Arrested is a message, destined to Ministry,
// indicating that wizard (wizardID)
// who start the attack (attackID)
// has been put in prison by Azkaban
type Arrested struct {
	ID            uuid.UUID `json:"id"`
	AttackID	  uuid.UUID `json:"attack_id"`
	WizardID      uuid.UUID `json:"wizard_id"`
	Message       string    `json:"message"`
}