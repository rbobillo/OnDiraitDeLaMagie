package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/hogwartsinventory"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/internal"
	"net/http"
)

// GetStudents function requests the Hogwarts Inventory
// to find every students
func GetStudents(w *http.ResponseWriter, db *sql.DB) error {
	internal.Debug(fmt.Sprintf("/wizards/"))

	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	query := "SELECT * FROM students"

	wizards, err := hogwartsinventory.GetAllStudents(db, query)

	if err == internal.ErrStudentsNotFounds {
		(*w).WriteHeader(http.StatusNotFound)
		internal.Warn("students doesn't exists")
		return err
	}
	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		internal.Warn("cannot find all students")
		return err
	}

	js, err := json.Marshal(wizards)

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn(fmt.Sprintf("cannot serialize students to JSON"))
		return err
	}

	_, err = fmt.Fprintf(*w, string(js))

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("cannot convert Body to JSON")
		return err
	}

	internal.Debug("all students has been found")

	return nil
}
