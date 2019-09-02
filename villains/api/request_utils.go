package api

import (
	"encoding/json"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/villains/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/villains/internal"
	"net/http"
)

// SingleVillainResponse try to serialize a villains to JSON
// then copy the JSON to the response header body
func SingleVillainResponse(villain dao.Villain, w *http.ResponseWriter) error {
	js, err := json.Marshal(villain)

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("cannot serialize Wizard to JSON")
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
