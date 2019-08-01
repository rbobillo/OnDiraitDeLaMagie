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
	"log"
	"net/http"

	_ "github.com/lib/pq" // go get -u github.com/lib/pq
)

// GetWizard function requests the Magic Inventory
// to find a specific wizard
// returns { "wizard" : <some_wizard> }
func GetWizard(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	id := mux.Vars(r)["id"]

	if len(mux.Vars(r)) == 0 {
		log.Print(GetWizards(w, db))
		return GetWizards(w, db)
	}

	log.Printf("/wizards/%s", id)
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	query := "SELECT * FROM wizards WHERE id = $1"

	wz, err := magicinventory.GetWizardsByID(db, query, id)

	if err == internal.ErrWizardsNotFounds {
		(*w).WriteHeader(http.StatusNotFound)
		return err
	}
	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		log.Printf("error: cannot find wizard %s", id)
	}

	js, err := json.Marshal(wz)

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		log.Fatal("error: cannot serialize Wizard to JSON")
		return err
	}
	_, err = fmt.Fprintf(*w, string(js))

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		return err
	}
	return err
}

// GetWizards function requests the Magic Inventory
// to find every wizards
func GetWizards(w *http.ResponseWriter, db *sql.DB)  error {
	log.Println("/wizards")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	query := "SELECT * FROM wizards"

	wizards, err := magicinventory.GetAllWizards(db, query)

	if err == internal.ErrWizardsNotFounds {
		(*w).WriteHeader(http.StatusNotFound)
		return err
	}
	if err != nil{
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		log.Printf("error: cannot find all wizards")
		return err
	}

	js, _ := json.Marshal(wizards)
	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		log.Fatal("error: cannot serialize Wizard to JSON")
		return err
	}

	_, err = fmt.Fprintf(*w, string(js))

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		return err
	}

	return nil
}
