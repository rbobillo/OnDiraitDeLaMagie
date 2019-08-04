package api

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/internal"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/magicinventory"
	"net/http"
)

// JailWizard put one wizard in jail by updating his jail status.
// TODO: better error handling (row not modified, db not reachable) - switch for err handling
func JailWizard(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	id := mux.Vars(r)["id"]

	internal.Debug(fmt.Sprintf("/wizards/%s/jail", id))

	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	query := "UPDATE wizards SET arrested = $2 WHERE id = $1 RETURNING *;"
	wz, err := magicinventory.UpdateWizardsByID(db, id, query, true)

	if err == internal.ErrWizardsNotFounds {
		(*w).WriteHeader(http.StatusNotFound)
		internal.Error(fmt.Sprintf("wizard %s doesn't exists", id))
		return err
	}

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		internal.Error(fmt.Sprintf("cannot arrest wizard %s", id))
		return err
	}

	err = SingleWizardResponse(wz, w)
	if  err != nil{
		return err
	}
	internal.Debug(fmt.Sprintf("wizard %s id in jail", id))

	return nil
}
