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

// SpawnWizard function requests the Magic Inventory to create a new wizard
// TODO: handle error on db error (with proper http return codes)
func SpawnWizard(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	var wizard dao.Wizard

	log.Println("/wizards/spawn")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&wizard)
	if err != nil {
		(*w).WriteHeader(http.StatusMethodNotAllowed)
		log.Println("warning: cannot convert Body to JSON")
		return err
	}

	err = magicinventory.CreateWizards(wizard, db)
	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		log.Println("warning: cannot insert new Wizard")
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

	(*w).WriteHeader(http.StatusCreated)
	log.Println("new wizard created")
	return nil
}
