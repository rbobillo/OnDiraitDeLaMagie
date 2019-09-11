package api

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/magic/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/magic/internal"
	"github.com/rbobillo/OnDiraitDeLaMagie/magic/magicinventory"
	"io"
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
	if err == io.EOF {
		internal.Debug("specified wizard not found. Generating random wizard")
		wizard, err = internal.GenerateSingleWizard()
	} else if err != nil {
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

	js, err := json.Marshal(wizard)
	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("cannot serialize Wizard to JSON")
		return err
	}

	err = postWizardToGuest(w, js)
	if err != nil{
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("error while requesting http://localhost:9092/guest/spawn")
		return err
	}

	internal.Info(fmt.Sprintf("a wizard spawned: %v", wizard))

	return nil
}

func postWizardToGuest(w *http.ResponseWriter, js []byte) (err error){
	req, err := http.NewRequest("POST", "http://localhost:9092/guest/spawn", bytes.NewBuffer(js))
	if req == nil || err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("error while creating post request")
		return err
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	defer resp.Body.Close()

	return err
}