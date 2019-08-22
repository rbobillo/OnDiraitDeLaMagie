package api

import (
	"database/sql"
	"encoding/json"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/dto"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/internal"
	uuid "github.com/satori/go.uuid"
	"net/http"
)
// AttackHogwarts stops Hogwarts activity
// while Hogwarts is not protected
// TODO: create DB insert in hogwartsinventory (actions table)
func AttackHogwarts(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	var attack dto.Attack

	internal.Info("/actions/attack : Hogwarts is under attack")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&attack)

	if err != nil {
		(*w).WriteHeader(http.StatusMethodNotAllowed)
		internal.Warn("warning: cannot convert Body to JSON")
		return err
	}
	// TODO: implement attack logic (impact on Hogwarts services + table insert...)

	help, err := json.Marshal(dto.Help{
		ID: uuid.Must(uuid.NewV4()),
		AttackID: attack.ID,
		Emergency: dto.Emergency{
			Quick: attack.Quick,
			Strong: attack.Strong,
		},
	})
	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("error: cannot serialize Wizard to JSON")
		return err
	}

	internal.Debug("alerting Ministry, Families, and Guest")
	// TODO: handle rabbit/queue disconnect errors ?

	internal.Publish("ministery", string(help))
	internal.Debug("Mail (alert) sent to ministry !")

	internal.Publish("families", string(help))
	internal.Debug("Mail (alert) sent to families") //TODO: better message

	internal.Publish("guest", string(help))
	internal.Debug("Mail (alert) sent to guest") //TODO: better message

	(*w).WriteHeader(http.StatusNoContent)
	return err
}