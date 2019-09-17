package dto

import uuid "github.com/satori/go.uuid"

// Arrest is the message from azkaban
// ordering Azkaban to put a wizard in jail
type Arrest struct {
	ID 		uuid.UUID `json:"id"`
	MagicID uuid.UUID `json:"magic_id"`
}