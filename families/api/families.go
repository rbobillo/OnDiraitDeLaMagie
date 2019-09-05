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

	wizard  :=  internal.FilterByID(wizards, id)

	if len(wizard) <= 0 {
		(*w).WriteHeader(http.StatusNotFound)
		internal.Warn(fmt.Sprintf("wizard %s doesn't exists", id))
		return err
	}

	//families, err := internal.FilterByFamilies(wizards, wizard[0].LastName)
	//if err != nil {
	//	(*w).WriteHeader(http.StatusNotFound)
	//	internal.Warn(fmt.Sprintf( "%s familly doesn't exists", wizard[0].LastName))
	//	return err
	//}
	//
	//js, err := json.Marshal(families)
	//
	//if err != nil {
	//	(*w).WriteHeader(http.StatusInternalServerError)
	//	internal.Warn(fmt.Sprintf("cannot serialize wizard to JSON"))
	//	return err
	//}
	//
	//_, err = fmt.Fprintf(*w, string(js))
	//
	//if err != nil {
	//	(*w).WriteHeader(http.StatusInternalServerError)
	//	internal.Warn("cannot convert Body to JSON")
	//	return err
	//}

	internal.Debug(fmt.Sprintf("the %s familly has been found", wizard[0].LastName))

	return err

}
