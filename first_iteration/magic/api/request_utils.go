package api

import (
	"encoding/json"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/internal"
	"net/http"
)

func SingleWizardResponse(wizard dao.Wizard, w *http.ResponseWriter) error{
	js, err := json.Marshal(wizard)

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Log(fmt.Sprintf("cannot serialize Wizard to JSON")).Error()
		return err
	}

	_, err = fmt.Fprintf(*w, string(js))

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Log(fmt.Sprintf("cannot convert Body to JSON")).Warn()
		return err
	}
	return nil
}
