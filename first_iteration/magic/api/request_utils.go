package api

import (
	"encoding/json"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/internal"
	"net/http"
)

// SingleWizardResponse try to serialize a wizard to JSON
// then copy the JSON to the response header body
func SingleWizardResponse(wizard dao.Wizard, w *http.ResponseWriter) error {
	js, err := json.Marshal(wizard)

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Error("cannot serialize Wizard to JSON")
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
