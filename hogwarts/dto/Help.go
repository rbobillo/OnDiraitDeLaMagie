package dto

import "github.com/satori/go.uuid" // go get github.com/satori/go.uuid

// Emergency defines the Help needed
type Emergency struct {
	Quick  bool `json:"quick"`
	Strong bool `json:"strong"`
}

// Help is the message sent by Hogwarts (via Owls)
// to the Ministry, asking for help
// Emergency is true if enemies are already here
// false, if they are on their way to Hogwarts
type Help struct {
	ID          uuid.UUID `json:"id"`
	AttackID    uuid.UUID `json:"attack_id"`
	HelpMessage string    `json:"help_message"`
	Emergency   Emergency `json:"emergency"`
}