package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rbobillo/OnDiraitDeLaMagie/families/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/families/internal"
	"io/ioutil"
	"net/http"
	"fmt"
)

// GetFamilies function requests the Magic Service
// to find a whole specific families
// returns { "families" : <some_families> }
func GetFamilies(w *http.ResponseWriter, r *http.Request) (err error) {
	id := mux.Vars(r)["id"]

	internal.Debug(fmt.Sprintf("/families/%s", id))

	url := "http://localhost:9090/wizards"

	resp, err := http.Get(url)
	if err != nil {
		(*w).WriteHeader(http.StatusNotFound)
		internal.Warn("failed to get /magic/wizards")
		return err
	}
	defer resp.Body.Close()


	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		(*w).WriteHeader(http.StatusNotFound)
		internal.Warn("failed to read response")
		return err
	}

	var wizards []dao.Wizard

	err = json.Unmarshal(body, &wizards)
	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("cannot unserialize JSON to wizards")
		return err
	}

	err, families := internal.Filter(wizards, "families", id)
	if err != nil {
		return err
	}

	_, err = json.Marshal(families)

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("cannot serialize Wizard to JSON")
		return err
	}

	(*w).WriteHeader(http.StatusCreated)

	internal.Debug(fmt.Sprintf("families of wizards %s has been found", id))

	return err

}
