package dto

import uuid "github.com/satori/go.uuid"

// Alert is the action of
// the Villains attacking Hogwarts
// quick/strong determine how efficient
// the attack will be
type Alert struct {
	ID           uuid.UUID `json:"id"`
	AttackID     uuid.UUID `json:"attack_id"`
	AlertMessage string    `json:"alert_message"`
}
