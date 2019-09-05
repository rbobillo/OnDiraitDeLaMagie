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
	"io/ioutil"
	"log"
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

	(*w).WriteHeader(http.StatusCreated)

	internal.Info(fmt.Sprintf("wizard has spawned: %v", wizard))
	req, err := http.NewRequest("POST", "http://localhost:9092/families/spawn", bytes.NewBuffer(js))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	if err != nil{
		internal.Warn("error while giving birth")
		log.Println(err)
		return err
	}
	log.Println(resp)
	return nil
}
