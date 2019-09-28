package api

import (
	"encoding/json"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/guest/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/guest/dto"
	"github.com/rbobillo/OnDiraitDeLaMagie/guest/internal"
	"github.com/rbobillo/OnDiraitDeLaMagie/guest/rabbit"
	uuid "github.com/satori/go.uuid"
	"net/http"
)
// SpawnNewSlot Assyncronly  the new wizard from magic
// then send mail to inform ministry
func StartSlotEvent(w *http.ResponseWriter, r *http.Request) (err error){

	//TODO: creat ascync process that call spawnNewSlot
	err = spawnNewSlot(w, r)
	if err != nil {
		internal.Warn("cannot ask Hogwarts for a new slot")
		return err
	}
	return err
}

func spawnNewSlot(w *http.ResponseWriter, r *http.Request) (err error){
	var guest dao.Wizard

	internal.Info("/guest/spawn : a new wizard want to visit Hogwarts!")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&guest)


	if err != nil {
		(*w).WriteHeader(http.StatusMethodNotAllowed)
		internal.Warn("cannot convert Body to JSON")
		return err
	}

	err = sendNewSlotRequest(guest)
	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		return err
	}

	_, err = json.Marshal(guest)

	if err != nil {
		(*w).WriteHeader(http.StatusInternalServerError)
		internal.Warn("cannot serialize Slot to JSON")
		return err
	}
	(*w).WriteHeader(http.StatusCreated)
	return err
}

func sendNewSlotRequest(guest dao.Wizard) (err error) {
	internal.Debug("Sending owls to inform Hogwarts...")

	slotRequest, err := json.Marshal(dto.Slot{
		ID : uuid.Must(uuid.NewV4()),
		WizardID: guest.ID,
		Message: fmt.Sprintf("%s %s want to visit Hogwarts", guest.FirstName, guest.LastName),
	})
	if err != nil {
		internal.Warn("cannot serialize mail(slot) to JSON")
		return err
	}

	rabbit.Publish("hogwarts", string(slotRequest))
	internal.Debug("Mail (slot) sent to hogwarts") //TODO: better message

	//// TODO: handle rabbit/queue disconnect errors ?
	return err
}
