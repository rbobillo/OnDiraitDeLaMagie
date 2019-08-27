package dao

import "github.com/satori/go.uuid" // go get github.com/satori/go.uuid

// Wizard is the content for the Magic Inventory DB
type Wizard struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Age       float64   `json:"age"`
	Category  string    `json:"category"` // Families, Guests, Villains
	Arrested  bool      `json:"arrested"`
	Dead      bool      `json:"dead"`
}