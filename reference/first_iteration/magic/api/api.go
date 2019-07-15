package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"OnDiraitDeLaMagie/reference/first_iteration/magic/dao"

	_ "github.com/lib/pq" // go get -u github.com/lib/pq
)

type Wizard = dao.Wizard

// GetWizards function requests the Magic Inventory
// to find wizards
func GetWizards(db *sql.DB, w *http.ResponseWriter) error {
	log.Println("/GetWizards")

	rows, err := db.Query("SELECT * FROM wizards")

	if err != nil {
		panic(err)
	}

	var wizards []Wizard

	for rows.Next() {
		var wz Wizard
		err = rows.Scan(&wz.ID, &wz.FirstName, &wz.LastName, &wz.Age, &wz.Category, &wz.Arrested, &wz.Dead)

		if err != nil {
			panic(err)
		}

		wizards = append(wizards, wz)
	}

	js, _ := json.Marshal(wizards)

	_, err = fmt.Fprintf(*w, string(js))

	return err
}

// Index function exposes the swagger API description
func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("/Index")
	fmt.Fprintf(w, "TODO: add Swagger API documentation")
}

// InitMagic starts the Magic service
func InitMagic(db *sql.DB) (err error) {
	http.HandleFunc("/", Index)

	http.HandleFunc("/wizards", func(w http.ResponseWriter, r *http.Request) {
		err = GetWizards(db, &w)
	})

	return err
}
