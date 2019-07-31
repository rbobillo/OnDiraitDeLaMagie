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

// KillWizard function request the Magic Inventory to update one wizard
func KillWizard(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	id := mux.Vars(r)["id"]

	log.Printf("/wizards/%s/die", id)

	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	query := "UPDATE wizards SET dead = $2 WHERE id = $1 RETURNING *;"
	wz, err := magicinventory.UpdateWizardsByID(db, id, query, true)

	if err == sql.ErrNoRows {
		(*w).WriteHeader(http.StatusNotFound)
		log.Println(fmt.Sprintf("wizard %s doesn't exists", id))
		return err
	}

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		log.Printf("error: cannot kill wizard %s", id)
		return err
	}

	js, err := json.Marshal(wz)

	// TODO: handle error

	_, err = fmt.Fprintf(*w, string(js))

	// TODO: handle error

	return nil
}
