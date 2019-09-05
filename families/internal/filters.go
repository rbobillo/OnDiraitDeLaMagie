package internal

import (
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/families/dao"
	uuid "github.com/satori/go.uuid"
)

type WizardType struct {
	ID       uuid.UUID
	strID    string
	LastName string
}


func Filter(wizards []dao.Wizard,  toCompare WizardType, f func (wizard dao.Wizard, toCompare WizardType) bool )  (fWizards []dao.Wizard) {

	for _, x := range wizards {
		if f(x, toCompare) {
			fWizards = append(fWizards, x)
		}
	}
	return fWizards
}

func FilterByID(wizards []dao.Wizard, id string) []dao.Wizard {

	isRightID := func(sWizard dao.Wizard,  toCompare WizardType) bool { return fmt.Sprintf("%v", sWizard.ID) == toCompare.strID }

	var toCompare WizardType
	toCompare.strID = id

	return Filter(wizards, toCompare, isRightID)
}
