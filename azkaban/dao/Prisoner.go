package dao

import (
	uuid "github.com/satori/go.uuid"
	"time"
)
// Prisoner is the content for the Azkaban Inventory DB
type Prisoner struct {
	ID       uuid.UUID `json:"id"`
	WizardID uuid.UUID `json:"wizard_id"`
	Arrest 	 time.Time `json:"arrest"`
}
