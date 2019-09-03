package api

import (
	"database/sql"
	"encoding/json"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/dto"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/hogwartsinventory"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/internal"
	uuid "github.com/satori/go.uuid"
	"net/http"
)
// AttackHogwarts stops Hogwarts activity
// while Hogwarts is not protected
func AttackHogwarts(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	var attack dao.Action

	internal.Info("/actions/attack : Hogwarts is under attack")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&attack)

	if err != nil {
		(*w).WriteHeader(http.StatusMethodNotAllowed)
		internal.Warn("cannot convert Body to JSON")
		return err
	}

	//TODO : Stop hogwarts prossess and w8 for the protection
	err = hogwartsinventory.CreateAttack(attack, db)
	if err != nil {
		(*w).WriteHeader(http.StatusUnprocessableEntity)
		internal.Warn("cannot insert new Attack")
		return err
	}

	err = sendAlertOwls(attack)
	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		return err
	}

	_, err = json.Marshal(attack)

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("cannot serialize Attack to JSON")
		return err
	}
	(*w).WriteHeader(http.StatusCreated)
	return err
}

func sendAlertOwls(attack dao.Action) (err error) {
	internal.Debug("Alerting Families, and Guest")

	alert, err := json.Marshal(dto.Alert{
		ID: uuid.Must(uuid.NewV4()),
		AttackID: attack.ID,
		Message: "Hogwarts is under attack",
	})
	if err != nil {
		internal.Warn("cannot serialize Attack to JSON")
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