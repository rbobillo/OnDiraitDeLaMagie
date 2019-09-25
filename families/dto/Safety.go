package dto

import "github.com/satori/go.uuid" // go get github.com/satori/go.uuid

// Safety is the message sent by Hogwarts (via Owls)
// to families, to inform that Hogwarts is
// no longer under attack
type Safety struct {
	ID            uuid.UUID `json:"id"`
	AttackID      uuid.UUID `json:"attack_id"`
	SafetyMessage string    `json:"safety_message"`
}