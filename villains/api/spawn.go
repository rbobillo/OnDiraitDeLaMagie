package api

import (
	"encoding/json"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/magic/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/villains/internal"
	"io/ioutil"
	"net/http"
)

//SpawnVillain function request the Magic Service
// to launch an attack on it randomly
//TODO : add random call.
func SpawnVillain(w *http.ResponseWriter, r *http.Request) (err error){
	internal.Debug("/villains/spawn")

	url := "http://localhost:9091/hogwarts/attack"

	resp, err := http.Get(url)

	if err != nil {
		(*w).WriteHeader(http.StatusNoContent)
		internal.Warn("fail to get /hogwarts/attack")
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		(*w).WriteHeader(http.StatusNotFound)
		internal.Warn("failed to read response")
		return err
	}

	var wizard dao.Wizard

	err = json.Unmarshal(body, &wizard)
	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("cannot unserialize JSON to wizard")
		return err
	}

	js, err := json.Marshal(wizard)

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("cannot serialize Wizard to JSON")
		return err
	}

	_, err = fmt.Fprintf(*w, string(js))

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("cannot convert Body to JSON")
		return err
	}
	internal.Debug("wizard has been found")

	return err
}