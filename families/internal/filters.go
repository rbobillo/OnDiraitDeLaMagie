package internal

import (
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/families/dao"
)
type FilterSet func(wizard dao.Wizard, arg	interface{}) bool
var filters = map[string]FilterSet{
	"families": filterByFamilies,
	"id"	  : filterByID,
	"age"	  :	filterByAge,
}

func filterByFamilies(wizard dao.Wizard, lastName interface{}) bool {
	if wizard.LastName != lastName {
		return false
	}
	return true
}

func filterByID(wizard dao.Wizard, wizardID interface{}) bool {

	id := fmt.Sprintf("%v", wizard.ID)
	wzID := fmt.Sprintf("%v", wizardID)

	if wzID != id{
		return false
	}
	return true
}

func filterByAge(wizard dao.Wizard, age interface{}) bool {
	if wizard.Age != age {
		return false
	}
	return true
}


func Filter(wizards []dao.Wizard, data string, value  interface{}) (err error, filteredWizard []dao.Wizard) {

	for _, entities := range wizards {
		if filters[data](entities, value) {
			filteredWizard = append(filteredWizard, entities)
		}
	}
	if len(filteredWizard) <= 0 {
		return fmt.Errorf("No match found"), filteredWizard
	}
	return nil, filteredWizard
}
