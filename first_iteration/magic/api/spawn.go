package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/magicinventory"
	"log"
	"net/http"
)

// SpawnWizard function requests the Magic Inventory
// to create a new wizard
// TODO: handle error on db error (duplicates, whatever...)
// returns { "wizard" : <some_wizard> }
func SpawnWizard(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	var wizard dao.Wizard

	log.Println("/wizards/spawn")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	// TODO: handle string to uuid unmarshal
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&wizard)

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		return err
	}

	err = magicinventory.CreateWizard(wizard, db)

	js, _ := json.Marshal(wizard)

	_, err = fmt.Fprintf(*w, string(js))

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		return err
	}

	return nil
}
