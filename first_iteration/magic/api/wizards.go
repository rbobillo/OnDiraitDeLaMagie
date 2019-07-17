// Package api is the entry point of 'magic' service
//
package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/dao"
	"log"
	"net/http"

	_ "github.com/lib/pq" // go get -u github.com/lib/pq
)

// GetWizardsByID function requests the Magic Inventory
// to find a specific wizard
// returns { "wizard" : <some_wizard> }
func GetWizardsByID(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	id := mux.Vars(r)["id"]

	if len(mux.Vars(r)) == 0 {
		return GetWizards(w, db)
	}

	log.Printf("/wizards/%s", id)
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	row := db.QueryRow("SELECT * FROM wizards WHERE id = $1", id)

	var wz dao.Wizard

	err = row.Scan(&wz.ID, &wz.FirstName, &wz.LastName, &wz.Age, &wz.Category, &wz.Arrested, &wz.Dead)

	if err != nil {
		(*w).WriteHeader(http.StatusNotFound)
		return err
	}

	js, _ := json.Marshal(map[string]dao.Wizard{"wizard": wz})

	_, err = fmt.Fprintf(*w, string(js))

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		return err
	}

	return nil
}

// GetWizards function requests the Magic Inventory
// to find every wizards
// returns { "wizards" : [ <every_wizards> ] }
func GetWizards(w *http.ResponseWriter, db *sql.DB) error {
	log.Println("/wizards")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	rows, err := db.Query("SELECT * FROM wizards")

	if err != nil {
		(*w).WriteHeader(http.StatusNotFound)
		return err
	}

	var wizards []dao.Wizard

	for rows.Next() {
		var wz dao.Wizard
		err = rows.Scan(&wz.ID, &wz.FirstName, &wz.LastName, &wz.Age, &wz.Category, &wz.Arrested, &wz.Dead)

		if err != nil {
			(*w).WriteHeader(http.StatusNotFound)
			return err
		}

		wizards = append(wizards, wz)
	}

	js, _ := json.Marshal(map[string][]dao.Wizard{"wizards": wizards})

	_, err = fmt.Fprintf(*w, string(js))

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		return err
	}

	return nil
}