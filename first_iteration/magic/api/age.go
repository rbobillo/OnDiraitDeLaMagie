package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/magicinventory"
	"log"
	"net/http"
)

// UpdateWizardsAges function request the Magic Inventory
// to update every wizard age by increment it n times
// Todo: Change n to json
func UpdateWizardsAge(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error){
	value := mux.Vars(r)["age"]

	log.Println("/wizards/age")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	err = magicinventory.UpdateWizards(db, "age", value)
	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		log.Println("error: cannot update wizards's age")
		return err
	}

	return nil
}

