package api

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/internal"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/magicinventory"
	"net/http"
)

// KillWizard function request the Magic Inventory to update one wizard
func KillWizard(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	id := mux.Vars(r)["id"]

	internal.Debug(fmt.Sprintf("/wizards/%s/die"), is)

	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	query := "UPDATE wizards SET dead = $2 WHERE id = $1 RETURNING *;"
	wizard, err := magicinventory.UpdateWizardsByID(db, id, query, true)

	if err == internal.ErrWizardsNotFounds {
		(*w).WriteHeader(http.StatusNotFound)
		internal.Error(fmt.Sprintf("wizard %s doesn't exists", id))
		return err
	}

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		internal.Error(fmt.Sprintf("cannot kill wizard %s", id))
		return err
	}

	err = SingleWizardResponse(wizard, w)
	if  err != nil{
		return err
	}

	internal.Info(fmt.Sprintf("wizard %s is dead", id))

	return nil
}
