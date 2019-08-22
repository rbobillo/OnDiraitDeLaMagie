package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/dto"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/internal"
	"net/http"
)

// ProtectHogwarts function cancels or avoids
// villains attack on Hogwarts
// TODO: create DB update in hogwartsinventory (actions table)
func ProtectHogwarts(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error){
	var protection dto.Protection

	id := mux.Vars(r)["id"]

	internal.Debug(fmt.Sprintf("/action/%s/protect : Ministery is protecting Hogwarts", id))
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&protection)
	if err != nil {
		(*w).WriteHeader(http.StatusMethodNotAllowed)
		internal.Warn("cannot convert Body to JSON")
		return err
	}

	// TODO: implement protection logic (simple table update ? -> Action attack status : stopped)

	internal.Debug("telling Famillies and Guest that Hogwarts is now safe")

	internal.Publish("families", "Hogwarts is safe") //Todo: better message
	internal.Debug("Mail (safety) sent to Families")

	internal.Publish("guest", "Hogwarts is now safe") //Todo:better message
	internal.Debug("Mail (safety) sent to Guest")

	(*w).WriteHeader(http.StatusNoContent)

	return err


}