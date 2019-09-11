package internal

import (
	"github.com/rbobillo/OnDiraitDeLaMagie/guest/dao"
)

func FilterWizards(wizards []dao.Wizard, f func (wizard dao.Wizard) bool )  (fWizards []dao.Wizard) {

	for _, x := range wizards {
		if f(x) {
			fWizards = append(fWizards, x)
		}
	}
	return fWizards
}

