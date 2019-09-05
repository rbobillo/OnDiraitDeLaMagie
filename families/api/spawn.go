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
	internal.Info("/actions/attack : Hogwarts is under attack")
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
	internal.Debug("new wizard just born ! Sending owls to inform ministry")

	alert, err := json.Marshal(dto.Birth{
		ID : uuid.Must(uuid.NewV4()),
		WizardID: wizard.ID,
		Message: fmt.Sprintf("%s %s just born !", wizard.FirstName, wizard.LastName),
	})
	if err != nil {
		internal.Warn("cannot serialize mail(bith) to JSON")
		return err
	}

	internal.Publish("families", string(alert))
	internal.Debug("Mail (alert) sent to families") //TODO: better message

	internal.Publish("guest", string(alert))
	internal.Debug("Mail (alert) sent to guest") //TODO: better message


	internal.Debug("Asking for help to Ministry")
	help, err := json.Marshal(dto.Help{
		ID: uuid.Must(uuid.NewV4()),
		AttackID: attack.ID,
		Message: "Hogwarts is under attack! Please send help",
		Emergency: dto.Emergency{
			Quick: true,
			Strong: true,
		},
	})
	if err != nil {
		internal.Warn("cannot serialize Attack to JSON")
		return err
	}

	internal.Publish("ministry", string(help))
	internal.Debug("Mail (help) sent to ministry !")

	// TODO: handle rabbit/queue disconnect errors ?
	return err
}
