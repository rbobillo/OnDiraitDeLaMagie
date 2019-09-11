package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rbobillo/OnDiraitDeLaMagie/guest/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/guest/internal"
	"io/ioutil"
	"log"
	"net/http"
	"fmt"
)

// GetGuest function requests the Magic Service
// to find a whole specific guest
// returns { "guest" : <some_guest> }
func GetGuestByID(w *http.ResponseWriter, r *http.Request) (err error) {
	id := mux.Vars(r)["id"]

	internal.Debug(fmt.Sprintf("/guest/%s", id))

	url := fmt.Sprintf("http://localhost:9090/wizards/%s", id)

	resp, err := http.Get(url)
	if err != nil {
		(*w).WriteHeader(http.StatusNotFound)
		log.Println(err)
		internal.Warn(fmt.Sprintf("failed to get /magic/wizards/{%s}", id))
		return err
	}
	defer resp.Body.Close()


	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		(*w).WriteHeader(http.StatusNotFound)
		internal.Warn("failed to read response")
		return err
	}

	var guest dao.Wizard

	err = json.Unmarshal(body, &guest)
	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("cannot unserialize JSON to wizards")
		return err
	}

	js, err := json.Marshal(guest)

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn(fmt.Sprintf("cannot serialize guest to JSON"))
		return err
	}

	_, err = fmt.Fprintf(*w, string(js))

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("cannot convert Body to JSON")
		return err
	}

	internal.Debug(fmt.Sprintf("the guest %s has been found", id))

	return err

}
