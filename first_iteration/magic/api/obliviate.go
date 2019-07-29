package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/magicinventory"
	"log"
	"net/http"
)

// ObliviateWizardByID obliviate a wizard from the magic
func ObliviateWizardByID(w *http.ResponseWriter, r *http.Request, db *sql.DB) (status error) {
	id := mux.Vars(r)["id"]

	log.Printf("/wizards/%s/obliviate", id)

	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	err := magicinventory.DeleteWizardsByID(db, id)

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		log.Printf("error: cannot obliviate wizards %s", id)
		return err
	}
	return nil
}
