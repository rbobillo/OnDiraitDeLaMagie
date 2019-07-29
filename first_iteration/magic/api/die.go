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

// KillWizardByID function request the Magic Inventory to update one wizard
func KillWizardByID(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	var wz dao.Wizard
	id := mux.Vars(r)["id"]

	log.Printf("/wizards/%s/die", id)

	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	query := fmt.Sprintf("UPDATE wizards SET dead = %t WHERE id = $1 AND dead != %t RETURNING *;", true, true)
	err = magicinventory.UpdateWizardsByID(db, query, id, wz)

	if err == sql.ErrNoRows {
		(*w).WriteHeader(http.StatusNotFound)
		log.Println(fmt.Sprintf("Wizard %s is already dead or doesn't exists", id))
		return err
	}

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		log.Printf("error: cannot kill wizards %s", id)
		return err
	}

	js, _ := json.Marshal(wz)
	_, err = fmt.Fprintf(*w, string(js))
	(*w).WriteHeader(http.StatusOK)
	return nil
}
