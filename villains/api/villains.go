// Package api is the entry point of 'villains' service
//
package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rbobillo/OnDiraitDeLaMagie/villains/internal"
	"github.com/rbobillo/OnDiraitDeLaMagie/villains/dao"
	"io/ioutil"
	"net/http"
	"fmt"
)

// GetVillain function requests the Magic Service
// to find a specific villain
// returns { "villain" : <some_villain> }
func GetVillain(w *http.ResponseWriter, r *http.Request) error {
	id := mux.Vars(r)["id"]
	internal.Debug(fmt.Sprintf("/villains/%s", id))

	url := fmt.Sprintf("http://localhost:9090/wizards/{%s}", id)

	resp, err := http.Get(url)
	if err != nil {
		(*w).WriteHeader(http.StatusNotFound)
		internal.Warn("failed to get /magic/wizard/{id}")
		return err
	}
	defer resp.Body.Close()


	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		(*w).WriteHeader(http.StatusNotFound)
		internal.Warn("failed to read response")
		return err
	}

	var villain dao.Villain

	err = json.Unmarshal(body, &villain)
	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("cannot unserialize JSON to villains")
		return err
	}

	err = SingleVillainResponse(villain, w)
	if err != nil {
		return err
	}

	internal.Debug(fmt.Sprintf("villains %s has been found", id))

	return err
}
