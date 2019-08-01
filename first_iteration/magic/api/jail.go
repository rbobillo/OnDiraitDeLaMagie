package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/internal"
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

	query := "UPDATE wizards SET arrested = $2 WHERE id = $1 RETURNING *;"
	wz, err := magicinventory.UpdateWizardsByID(db, id, query, true)

	if err == internal.ErrWizardsNotFounds {
		(*w).WriteHeader(http.StatusNotFound)
		log.Println(fmt.Sprintf("wizard %s doesn't exists", id))
		return err
	}

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		log.Printf("error: cannot arrest wizard %s", id)
		return err
	}

	js, err := json.Marshal(wz)

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		log.Fatal("error: cannot serialize Wizard to JSON")
		return err
	}

	_, err = fmt.Fprintf(*w, string(js))

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		log.Fatal("warning: cannot convert Body to JSON")
		return err
	}

	return nil
}
