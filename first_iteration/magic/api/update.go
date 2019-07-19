package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/magicinventory"
	"log"
	"net/http"
)

// UpdateWizardsAges function request the Magic Inventory
// to update every wizard age by increment it n times
func UpdateWizardsAge(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error){
	var n float64 = 2
	var wizard dao.Wizard

	log.Println("/wizards/updateAge")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&wizard)

	if err != nil {
		(*w).WriteHeader(http.StatusMethodNotAllowed)
		log.Println("warning: cannot convert Body to JSON")
		return err
	}

	err = magicinventory.UpdateWizard(wizard, db, n)

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		log.Println("error: cannot update wizards's age")
		return err
	}

	js, err := json.Marshal(wizard)

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		log.Fatal("error: cannot serialize Wizard to JSON")
		return err
	}
	_, err = fmt.Fprintf(*w, string(js))
	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		log.Fatal("warning: cannot convert Body to JSON")
		return err
	}
	return nil
}