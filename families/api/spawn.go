package api

import (
	"encoding/json"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/families/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/families/dto"
	"github.com/rbobillo/OnDiraitDeLaMagie/families/internal"
	uuid "github.com/satori/go.uuid"
	"net/http"
)

func SpawnNewBorn(w *http.ResponseWriter, r *http.Request) (err error){
	var born dao.Wizard
	internal.Info("/families/spawn : a new wizard just born !")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&born)

	if err != nil {
		(*w).WriteHeader(http.StatusMethodNotAllowed)
		internal.Warn("cannot convert Body to JSON")
		return err
	}

	err = sendBirthOwls(born)
	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		return err
	}

	_, err = json.Marshal(born)

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("cannot serialize Attack to JSON")
		return err
	}
	(*w).WriteHeader(http.StatusCreated)
	return err
}

func sendBirthOwls(wizard dao.Wizard) (err error) {
	internal.Debug("Sending owls to inform ministry")

	alert, err := json.Marshal(dto.Birth{
		ID : uuid.Must(uuid.NewV4()),
		WizardID: wizard.ID,
		Message: fmt.Sprintf("%s %s just born !", wizard.FirstName, wizard.LastName),
	})
	if err != nil {
		internal.Warn("cannot serialize mail(birth) to JSON")
		return err
	}

	internal.Publish("ministry", string(alert))
	internal.Debug("Mail (birth) sent to ministry") //TODO: better message

	// TODO: handle rabbit/queue disconnect errors ?
	return err
}
