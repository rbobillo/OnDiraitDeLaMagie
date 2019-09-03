package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/hogwartsinventory"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/internal"
	"net/http"
)

// GetStudent function requests the Magic Inventory
// to find a specific wizard
// returns { "wizard" : <some_wizard> }
func GetStudent(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	id := mux.Vars(r)["id"]

	internal.Debug(fmt.Sprintf("/students/%s", id))

	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	query := "SELECT * FROM wizards WHERE id = $1"

	student, err := hogwartsinventory.GetStudentByID(db, query, id)

	if err == internal.ErrStudentsNotFounds {
		(*w).WriteHeader(http.StatusNotFound)
		internal.Warn(fmt.Sprintf("student %s may not exist", id))
		return err
	}
	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		internal.Warn(fmt.Sprintf("cannot find wizard %s (%s)", id, err))
		return err
	}

	err = SingleStudentResponse(student, w)
	if err != nil {
		return err
	}
	internal.Debug(fmt.Sprintf("wizard %s hase been found", id))

	return err
}
// GetStudents function requests the Hogwarts Inventory
// to find every students
func GetStudents(w *http.ResponseWriter, db *sql.DB) error {
	internal.Debug(fmt.Sprintf("/students/"))

	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	query := "SELECT * FROM students"

	students, err := hogwartsinventory.GetAllStudents(db, query)

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

	js, err := json.Marshal(students)

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
