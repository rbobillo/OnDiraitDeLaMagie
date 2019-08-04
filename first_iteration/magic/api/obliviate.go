package api

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/internal"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/magicinventory"
	"net/http"
)

// ObliviateWizard obliviate a wizard from the magic
func ObliviateWizard(w *http.ResponseWriter, r *http.Request, db *sql.DB) error {
	id := mux.Vars(r)["id"]

	internal.Debug(fmt.Sprintf("/wizards/%s/obliviate", id))

	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	err := magicinventory.DeleteWizardsByID(db, id)

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		internal.Warn(fmt.Sprintf("cannot obliviate wizards %s", id))
		return err
	}

	(*w).WriteHeader(http.StatusNoContent)
	internal.Info(fmt.Sprintf("wizard %s have been oblivited", id))

	return nil
}
