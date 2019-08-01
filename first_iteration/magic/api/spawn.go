package api

import (
	"database/sql"
	"encoding/json"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/internal"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/magicinventory"
	"net/http"
	"fmt"
)

// SpawnWizard function requests the Magic Inventory to create a new wizard
// TODO: handle error on db error (with proper http return codes)
func SpawnWizard(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	var wizard dao.Wizard

	internal.Log(fmt.Sprintf("/wizards/spawn")).Debug()

	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&wizard)

	if err != nil {
		(*w).WriteHeader(http.StatusMethodNotAllowed)
		internal.Log(fmt.Sprintf("cannot convert Body to JSON")).Warn()
		return err
	}

	err = magicinventory.CreateWizards(wizard, db)

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		internal.Log(fmt.Sprintf("cannot insert new Wizard")).Warn()
		return err
	}

	_, err = json.Marshal(wizard)

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Log(fmt.Sprintf("cannot serialize Wizard to JSON")).Error()
		return err
	}

	(*w).WriteHeader(http.StatusCreated)

	internal.Log(fmt.Sprintf("new wizard created")).Debug()

	return nil
}
