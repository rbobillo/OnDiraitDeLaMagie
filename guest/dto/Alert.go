package dto

import uuid "github.com/satori/go.uuid"

// Alert is the message of
// sent by  Hogwarts
// to inform guest that
// Hogwarts is under attack
type Alert struct {
	ID        uuid.UUID   `json:"id"`
	AttackID  uuid.UUID   `json:"attack_id"`
	AlertMessage   string `json:"alert_message"`
}