package api

import (
	"database/sql"
	"encoding/json"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/hogwartsinventory"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/internal"
	"net/http"
)

func VisitHogwarts(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error){
	var visit dao.Action

	internal.Info("/actions/visit : Hogwarts receive a visit")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&visit)

	if err != nil {
		(*w).WriteHeader(http.StatusMethodNotAllowed)
		internal.Warn("cannot convert Body to JSON")
		return err
	}

	visit.Status = "ongoing"

	err = hogwartsinventory.CreateVisit(visit, db)
	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		internal.Warn("cannot insert new Visit")
		return err
	}

	_, err = json.Marshal(visit)

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("cannot serialize Visit to JSON")
		return err
	}

	(*w).WriteHeader(http.StatusCreated)

	return  err
}