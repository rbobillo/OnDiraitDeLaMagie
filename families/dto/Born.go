package dto

import "github.com/satori/go.uuid" // go get github.com/satori/go.uuid

// Born is the new wizard from the Magic Inventory DB
type Born struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Age       float64   `json:"age"`
	Category  string    `json:"category"` // Families, Guests, Villains
	Arrested  bool      `json:"arrested"`
	Dead      bool      `json:"dead"`
}

// Birth is the message sent by Families (via Owls)
// to the Ministry, announce a new born
// Wizards in the families
type Birth struct {
	ID        uuid.UUID `json:"id"`
	BornID	  uuid.UUID `json:"bornID"`
	Message   string    `json:"message"`
}