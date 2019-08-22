// Package hogwartsinventory is used to setup and manipulate
// the magic database (hogwartsinventory)
package hogwartsinventory

import (
	"database/sql"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/internal"
	log "github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/internal"

)

func InitHogwartsInventory(psqlInfo string) (*sql.DB, error){
	db, err := sql.Open("postgres", psqlInfo)
	internal.HandleError(err, "failed to open a sql connection to posgres", log.Warn)

	initActionsQuery :=
		`create table if not exists actions (
	id        uuid not    null primary key,
    wizard_id uuid not    null,
    category  varchar(50) not null,
    status    varchar(50) not null
	); alter table actions owner to hogwarts;`

	_, err = db.Query(initActionsQuery)
	internal.HandleError(err, "failed to create 'actions' table", log.Warn)

	internal.Info("actions")
	return db, err
}