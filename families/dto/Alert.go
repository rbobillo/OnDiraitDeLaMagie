package dto

import uuid "github.com/satori/go.uuid"

// Alert is the message of
// sent by  Hogwarts
// to inform families that
// Hogwarts is under attack
type Alert struct {
	ID        uuid.UUID `json:"id"`
	AttackID  uuid.UUID `json:"attack_id"`
	Message   string    `json:"message"`
}