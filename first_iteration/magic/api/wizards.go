// Package api is the entry point of 'magic' service
//
package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/magicinventory"
	"log"
	"net/http"

	_ "github.com/lib/pq" // go get -u github.com/lib/pq
)

// GetWizard function requests the Magic Inventory
// to find a specific wizard
// returns { "wizard" : <some_wizard> }
func GetWizard(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	var wz dao.Wizard
	id := mux.Vars(r)["id"]

	if len(mux.Vars(r)) == 0 {
		log.Print(GetWizards(w, db))
		return GetWizards(w, db)
	}

	log.Printf("/wizards/%s", id)
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	query := "SELECT * FROM wizards WHERE id = $1"

	err = magicinventory.GetWizardsByID(db, query, id, wz)

	js, _ := json.Marshal(wz)
	_, err = fmt.Fprintf(*w, string(js))

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		return err
	}

	(*w).WriteHeader(http.StatusOK)
	return nil
}

// GetWizards function requests the Magic Inventory
// to find every wizards
func GetWizards(w *http.ResponseWriter, db *sql.DB) error {
	var wizards []dao.Wizard
	log.Println("/wizards")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	query := "SELECT * FROM wizards"

	err := magicinventory.GetAllWizards(db, query, wizards)

	js, _ := json.Marshal(wizards)

	_, err = fmt.Fprintf(*w, string(js))

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		return err
	}

	(*w).WriteHeader(http.StatusOK)
	return nil
}
