package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/internal"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/magicinventory"
	"net/http"
	"fmt"
)

// ObliviateWizard obliviate a wizard from the magic
func ObliviateWizard(w *http.ResponseWriter, r *http.Request, db *sql.DB) error {
	id := mux.Vars(r)["id"]

	internal.Log(fmt.Sprintf("/wizards/%s/obliviate", id)).Debug()

	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	err := magicinventory.DeleteWizardsByID(db, id)

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		internal.Log(fmt.Sprintf("cannot obliviate wizards %s", id)).Error()
		return err
	}

	internal.Log(fmt.Sprintf("Wizard %s have been oblivited", id)).Debug()

	return nil
}
