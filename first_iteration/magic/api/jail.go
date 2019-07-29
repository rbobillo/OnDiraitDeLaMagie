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
)

// JailWizardByID put one wizard in jail by updating his jail status.
func JailWizardByID(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	var wz dao.Wizard
	id := mux.Vars(r)["id"]

	log.Printf("/wizards/%s/jail", id)

	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	query := fmt.Sprintf("UPDATE wizards SET arrested = %t WHERE id = $1 AND arrested != %t RETURNING *;", true, true)
	err = magicinventory.UpdateWizardsByID(db, query, id, wz)

	if err == sql.ErrNoRows {
		log.Println(fmt.Sprintf("Wizard %s is already in jail or wizard doesn't exists", id))
		return err
	}

	if err != nil {
		log.Printf("error: cannot arrest wizards %s", id)
		return err
	}
	js, _ := json.Marshal(wz)
	_, err = fmt.Fprintf(*w, string(js))
	(*w).WriteHeader(http.StatusOK)

	return nil
}
