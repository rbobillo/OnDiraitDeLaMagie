package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/magic/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/magic/internal"
	"github.com/rbobillo/OnDiraitDeLaMagie/magic/magicinventory"
	"net/http"
)

// SpawnWizard function requests the Magic Inventory to create a new wizard
// TODO: handle error on db error (with proper http return codes)
func SpawnWizard(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	var wizard dao.Wizard

	internal.Debug("/wizards/spawn")

	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&wizard)

	if err != nil {
		(*w).WriteHeader(http.StatusMethodNotAllowed)
		internal.Warn("cannot convert Body to JSON")
		return err
	}

	err = magicinventory.CreateWizards(wizard, db)

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		internal.Warn("cannot insert new Wizard")
		return err
	}

	_, err = json.Marshal(wizard)

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("cannot serialize Wizard to JSON")
		return err
	}

	(*w).WriteHeader(http.StatusCreated)

	internal.Info(fmt.Sprintf("wizard has spawned: %v", wizard))

	return nil
}
