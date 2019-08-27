package dao

import uuid "github.com/satori/go.uuid"

// Action is the content for the Hogwarts Inventory DB
type Action struct {
	ID         uuid.UUID `json:"id"`
	First_name string    `json:"first_name"`
	Last_name  string    `json:"last_name"`
	Age        string    `json:"age"`
	Arrested   string    `json:"arrested"`
	Category   string    `json:"category"`
	Dead       string    `json:"dead"`
}