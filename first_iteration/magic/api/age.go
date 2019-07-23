package api

import (
	"database/sql"
	"encoding/json"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/magicinventory"
	"log"
	"net/http"
)

// UpdateWizardsAges function request the Magic Inventory
// to update every wizard age by increment it n times
// Todo: Change n to json
func UpdateWizardsAge(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error){
	var wizard dao.Wizard

	log.Println("/wizards/age")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&wizard)

	if err != nil {
		(*w).WriteHeader(http.StatusMethodNotAllowed)
		log.Println("warning: cannot convert Body to JSON")
		return err
	}
	log.Println(wizard.Age)

	err = magicinventory.UpdateWizards(db, "age", wizard.Age)
	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		log.Println("error: cannot update wizards's age")
		return err
	}

	return nil
}

