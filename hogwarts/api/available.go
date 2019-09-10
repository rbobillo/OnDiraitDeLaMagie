package api
//
//import (
//	"database/sql"
//	"fmt"
//	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/dto"
//	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/hogwartsinventory"
//	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/internal"
//)
//
//// CheckSlot request the hogwarts inventory to find the number of ongoing visits
//func CheckSlot(slot dto.Slot, db *sql.DB) (err error, available int ){
//
//	query := "SELECT * FROM actions WHERE status = 'ongoing' and action = 'visit'"
//
//	ongoing, err := hogwartsinventory.GetActions(db, query)
//	if err !=  nil {
//		internal.Warn("cannot get actions in hogwarts inventory")
//		return err, 0
//	}
//	if len(ongoing) > 10 {
//		return fmt.Errorf("hogwarts have 10 visit ongoing"), 0
//	}
//	return err, 10
//}
