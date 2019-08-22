package dao

import uuid "github.com/satori/go.uuid"

// Wizard is the content for the Magic Inventory DB
type Wizard struct {
	ID         uuid.UUID `json:"id"`
	First_name string    `json:"first_name"`
	Last_name  string    `json:"last_name"`
	Age        string    `json:"age"`
	Arrested   string    `json:"arrested"`
	Category   string    `json:"category"`
	Dead       string    `json:"dead"`
}