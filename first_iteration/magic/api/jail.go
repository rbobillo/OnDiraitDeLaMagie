package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/magicinventory"
	"log"
	"net/http"
)

// JailWizard put one wizard in jail by updating his jail status.
// TODO: better error handling (row not modified, db not reachable) - switch for err handling
func JailWizard(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	id := mux.Vars(r)["id"]

	log.Printf("/wizards/%s/jail", id)

	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	query := fmt.Sprintf("UPDATE wizards SET arrested = %t WHERE id = $1 AND arrested != %t RETURNING *;", true, true)
	wz, err := magicinventory.UpdateWizardsByID(db, query, id)

	if err != nil {

		log.Printf("error: cannot arrest wizards %s", id)
		return err
	}

	js, _ := json.Marshal(wz)
	_, err = fmt.Fprintf(*w, string(js))

	return nil
}
