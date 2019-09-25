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
func LeaveHogwarts(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error){
	id := mux.Vars(r)["id"]

	internal.Debug(fmt.Sprintf("/actions/%s/leave", id))
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	query := "UPDATE actions SET status = $2 WHERE id = $1 RETURNING *;"
	act, err := hogwartsinventory.UpdateActionsByID(db, query ,id, "done")

	if err == internal.ErrActionsNotFounds {
		(*w).WriteHeader(http.StatusNotFound)
		internal.Warn(fmt.Sprintf("visits %s doesn't exists", id))
		return err
	}

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		internal.Warn(fmt.Sprintf("cannot stop visits %s", id))
		return err
	}

	err = SingleActionResponse(act, w)
	if err != nil {

		return err
	}
	internal.Info(fmt.Sprintf("visits %s was stopped", id))

	sendAvailableOwls()

	return err
}

func sendAvailableOwls(){
	internal.Debug("telling Guests that Hogwarts is now ready to start a new visits")

	rabbit.Publish("guest", "Hogwarts is now available") //Todo:better message
	internal.Debug("Mail (safety) sent to Guest")
}