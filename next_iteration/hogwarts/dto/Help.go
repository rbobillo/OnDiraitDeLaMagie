package dao

import "github.com/satori/go.uuid" // go get github.com/satori/go.uuid

// Help is the message sent by Hogwarts (via Owls)
// to the Ministry, asking for help
// Emergency is true if enemies are already here
// false, if they are on their way to Hogwarts
type Help struct {
	ID        uuid.UUID `json:"id"`
	Message   string    `json:"message"`
	Emergency bool      `json:"emergency"`
}
