// Package api is the entry point of 'magic' service
//
package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/dao"
	"log"
	"net/http"

	_ "github.com/lib/pq" // go get -u github.com/lib/pq
)

// GetWizards function requests the Magic Inventory
// to find wizards
func GetWizards(db *sql.DB, w *http.ResponseWriter) error {
	log.Println("/GetWizards")

	rows, err := db.Query("SELECT * FROM wizards")

	if err != nil {
		panic(err)
	}

	var wizards []dao.Wizard

	for rows.Next() {
		var wz dao.Wizard
		err = rows.Scan(&wz.ID, &wz.FirstName, &wz.LastName, &wz.Age, &wz.Category, &wz.Arrested, &wz.Dead)

		if err != nil {
			panic(err)
		}

		wizards = append(wizards, wz)
	}

	js, _ := json.Marshal(map[string][]dao.Wizard{"wizards": wizards})

	_, err = fmt.Fprintf(*w, string(js))

	return err
}

// Index function exposes the swagger API description
func Index(w *http.ResponseWriter, r *http.Request) (err error) {
	log.Println("/Index")
	_, err = fmt.Fprintf(w, "TODO: add Swagger API documentation")

	return err
}

// InitMagic starts the Magic service
func InitMagic(db *sql.DB) (err error) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err = Index(&w, r)
	})

	http.HandleFunc("/wizards", func(w http.ResponseWriter, r *http.Request) {
		err = GetWizards(db, &w)
	})

	return err
}
