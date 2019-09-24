package api

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/hogwartsinventory"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/internal"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/rabbit"
	"net/http"
)

// ProtectHogwarts function cancels or avoids
// villains attack on Hogwarts
func ProtectHogwarts(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error){
	id := mux.Vars(r)["id"]

	internal.Debug(fmt.Sprintf("/action/%s/protect : Ministery is protecting Hogwarts", id))
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	// TODO: implement protection logic (simple table update ? -> Action attack status : stopped)
	query := "UPDATE actions SET status = $2 WHERE id = $1 RETURNING *;"
	act, err := hogwartsinventory.UpdateActionsByID(db, query ,id, "done")

	if err == internal.ErrActionsNotFounds {
		(*w).WriteHeader(http.StatusNotFound)
		internal.Warn(fmt.Sprintf("attack %s doesn't exists", id))
		return err
	}

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		internal.Warn(fmt.Sprintf("cannot stop attack %s", id))
		return err
	}

	err = SingleActionResponse(act, w)
	if err != nil {

		return err
	}
	internal.Info(fmt.Sprintf("attack %s was stopped", id))

	sendSafetyOwls()

	return err
}

func sendSafetyOwls(){
	internal.Debug("telling Famillies and Guest that Hogwarts is now safe")

	rabbit.Publish("families", "Hogwarts is safe") //Todo: better message
	internal.Debug("Mail (safety) sent to Families")

	rabbit.Publish("guest", "Hogwarts is now safe") //Todo:better message
	internal.Debug("Mail (safety) sent to Guest")
}