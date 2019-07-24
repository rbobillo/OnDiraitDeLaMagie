package api

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/rbobillo/OnDiraitDeLaMagie/next_iteration/hogwarts/dto"
	"github.com/rbobillo/OnDiraitDeLaMagie/next_iteration/hogwarts/internal"
	"log"
	"net/http"
)

// ProtectHogwarts function cancels or avoids
// villains attack on Hogwarts
// TODO: create DB update in hogwartsinventory (actions table)
func ProtectHogwarts(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	var protection dto.Protection

	id := mux.Vars(r)["id"]

	log.Printf("/actions/%s/protect : Ministry is protecting Hogwarts", id)
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&protection)

	if err != nil {
		(*w).WriteHeader(http.StatusMethodNotAllowed)
		log.Println("warning: cannot convert Body to JSON")
		return err
	}

	// TODO: implement protection logic (simple table update ? -> Action attack status : stopped)

	log.Println("Telling Families and Guests that Hogwarts is now safe")

	internal.Publish("families", "Hogwarts is safe") // TODO: better message
	log.Println("Mail (safety) sent to Families")

	internal.Publish("guests", "Hogwarts is safe") // TODO: better message
	log.Println("Mail (safety) sent to Guests")

	(*w).WriteHeader(http.StatusNoContent)

	return err
}
