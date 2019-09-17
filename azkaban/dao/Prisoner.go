package dao

import (
	uuid "github.com/satori/go.uuid"
	"time"
)
// Prisoner is the content for the Azkaban Inventory DB
type Prisoner struct {
	ID      uuid.UUID `json:"id"`
	MagicID uuid.UUID `json:"magic_id`
	Arrest 	time.Time `json:"arrest"`
}
