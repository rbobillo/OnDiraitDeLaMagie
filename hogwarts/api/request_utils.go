package api

import (
	"encoding/json"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/internal"
	"net/http"
)

// SingleStudentResponse try to serialize a wizard to JSON
// then copy the JSON to the response header body
func SingleStudentResponse(student dao.Student, w *http.ResponseWriter) error {
	js, err := json.Marshal(student)

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("cannot serialize Student to JSON")
		return err
	}

	_, err = fmt.Fprintf(*w, string(js))

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("cannot convert Body to JSON")
		return err
	}
	return nil
}
