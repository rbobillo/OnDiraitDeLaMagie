package api
//
//import (
//	"encoding/json"
//	"github.com/rbobillo/OnDiraitDeLaMagie/families/internal"
//	"github.com/rbobillo/OnDiraitDeLaMagie/magic/magicinventory"
//	"io"
//	"net/http"
//)
//
//func SpawnNewBorn(w *http.ResponseWriter, r *http.Request) (err error){
//	var wizard dao.Wizard
//
//	internal.Debug("/wizards/spawn")
//
//	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")
//
//	decoder := json.NewDecoder(r.Body)
//	err = decoder.Decode(&wizard)
//	if err == io.EOF {
//		internal.Debug("specified wizard not found. Generating random wizard")
//		wizard, err = internal.GenerateSingleWizard()
//	} else if err != nil {
//		(*w).WriteHeader(http.StatusMethodNotAllowed)
//		internal.Warn("cannot convert Body to JSON")
//		return err
//	}
//
//	err = magicinventory.CreateWizards(wizard, db)
//
//	if err != nil {
//		(*w).WriteHeader(http.StatusUnprocessableEntity)
//		internal.Warn("cannot insert new Wizard")
//		return err
//	}
//
//	_, err = json.Marshal(wizard)
//
//	if err != nil {
//		(*w).WriteHeader(http.StatusInternalServerError)
//		internal.Warn("cannot serialize Wizard to JSON")
//		return err
//	}
//
//	(*w).WriteHeader(http.StatusCreated)
//
//	internal.Info(fmt.Sprintf("wizard has spawned: %v", wizard))
//
//	return nil
//}
