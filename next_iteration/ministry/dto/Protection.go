package dto

import "github.com/satori/go.uuid" // go get github.com/satori/go.uuid

// Protection is the action of
// the Ministry to help Hogwarts
// quick/strong determine how efficient
// the protection will be
type Protection struct {
	ID     uuid.UUID `json:"id"`
	Quick  bool      `json:"quick"`
	Strong bool      `json:"strong"`
}