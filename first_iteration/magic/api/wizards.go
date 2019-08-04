// Package api is the entry point of 'magic' service
//
package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/internal"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/magicinventory"
	"net/http"

	_ "github.com/lib/pq" // go get -u github.com/lib/pq
)

// GetWizard function requests the Magic Inventory
// to find a specific wizard
// returns { "wizard" : <some_wizard> }
func GetWizard(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	id := mux.Vars(r)["id"]

	internal.Debug(fmt.Sprintf("/wizards/%s", id))

	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	query := "SELECT * FROM wizards WHERE id = $1"

	wizard, err := magicinventory.GetWizardsByID(db, query, id)

	if err == internal.ErrWizardsNotFounds {
		(*w).WriteHeader(http.StatusNotFound)
		internal.Warn(fmt.Sprintf("wizard %s may not exist", id))
		return err
	}
	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		internal.Warn(fmt.Sprintf("cannot find wizard %s (%s)", id, err))
		return err
	}

	err = SingleWizardResponse(wizard, w)
	if err != nil {
		return err
	}
	internal.Debug(fmt.Sprintf("Wizard %s have bee found", id))

	return err
}

// GetWizards function requests the Magic Inventory
// to find every wizards
func GetWizards(w *http.ResponseWriter, db *sql.DB) error {
	internal.Debug(fmt.Sprintf("/wizards/"))

	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	query := "SELECT * FROM wizards"

	wizards, err := magicinventory.GetAllWizards(db, query)

	if err == internal.ErrWizardsNotFounds {
		(*w).WriteHeader(http.StatusNotFound)
		internal.Warn("wizards doesn't exists")
		return err
	}
	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		internal.Warn("cannot find all wizards")
		return err
	}

	js, err := json.Marshal(wizards)

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn(fmt.Sprintf("cannot serialize Wizard to JSON"))
		return err
	}

	_, err = fmt.Fprintf(*w, string(js))

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("cannot convert Body to JSON")
		return err
	}

	internal.Debug("all wizards have been found")

	return nil
}
