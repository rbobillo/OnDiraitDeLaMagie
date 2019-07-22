package api

import (
	"database/sql"
	"encoding/json"
	"github.com/rbobillo/OnDiraitDeLaMagie/next_iteration/hogwarts/dto"
	"github.com/rbobillo/OnDiraitDeLaMagie/next_iteration/hogwarts/internal"
	uuid "github.com/satori/go.uuid"
	"log"
	"net/http"
)

// AttackHogwarts stops Hogwarts activity
// while Hogwarts is not protected
// TODO: create status in hogwartsinventory ?
func AttackHogwarts(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	var attack dto.Attack

	log.Println("/attack : Hogwarts is under attack")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&attack)

	if err != nil {
		(*w).WriteHeader(http.StatusMethodNotAllowed)
		log.Println("warning: cannot convert Body to JSON")
		return err
	}

	// TODO: implement attack logic (impact on Hogwarts services...)

	help, err := json.Marshal(dto.Help{
		ID: uuid.Must(uuid.NewV4()),
		Message:"HELP",
		Emergency: dto.Emergency{
			Quick: attack.Quick,
			Strong: attack.Strong,
		},
	})

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		log.Fatal("error: cannot serialize Wizard to JSON")
		return err
	}

	log.Println("Alerting Ministry, Families and Guests")

	// TODO: handle rabbit/queue disconnect errors ?

	internal.Publish("ministry", string(help))
	log.Println("Mail (alert) sent to ministry !")

	internal.Publish("families", "Hogwarts is under attack") // TODO: better message
	log.Println("Mail (alert) sent to Families")

	internal.Publish("guests", "Hogwarts is under attack") // TODO: better message
	log.Println("Mail (alert) sent to Guests")

	(*w).WriteHeader(http.StatusNoContent)

	return err
}
