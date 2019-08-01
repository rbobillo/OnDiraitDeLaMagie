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

// AgeWizards function request the Magic Inventory to update every wizard age by increment it n times
func AgeWizards(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	var wizard dao.Wizard

	internal.Log(fmt.Sprintf("/wizards/age")).Debug()

	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&wizard)

	if err != nil {
		(*w).WriteHeader(http.StatusMethodNotAllowed)
		internal.Log(fmt.Sprintf("cannot convert Body to JSON")).Warn()
		return err
	}

	query := "UPDATE wizards SET age = age + $1;"
	err = magicinventory.UpdateWizards(db, query, wizard.Age)

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		internal.Log(fmt.Sprintf("cannot update wizards's age")).Error()
		return err
	}

	(*w).WriteHeader(http.StatusNoContent)
	internal.Log(fmt.Sprintf("wizards age successfully updated")).Debug()
	return nil
}
