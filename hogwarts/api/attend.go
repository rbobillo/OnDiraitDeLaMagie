package api


import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/hogwartsinventory"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/internal"
	"net/http"
	"fmt"
)

// AttendHogwarts function make
// a student attend Hogwarts
func AttendHogwarts(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error){
	id := mux.Vars(r)["id"]

	internal.Debug(fmt.Sprintf("/students/%s/attend", id))
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	query := "UPDATE students SET status = $2 WHERE id = $1 RETURNING *;"
	act, err := hogwartsinventory.UpdateStudentsByID(db, query ,id, "attending")

	if err == internal.ErrActionsNotFounds {
		(*w).WriteHeader(http.StatusNotFound)
		internal.Warn(fmt.Sprintf("student %s doesn't exists", id))
		return err
	}

	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		internal.Warn(fmt.Sprintf("student %s cannot attend Hogwarts", id))
		return err
	}

	err = SingleActionResponse(act, w)
		if err != nil {

		return err
	}
	internal.Info(fmt.Sprintf("student %s is attending hogwarts", id))

	return err
}