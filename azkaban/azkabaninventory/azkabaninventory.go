// Package azkabaninventory is used to setup and manipulate
// the azkaban database (azkabaninventory)
package azkabaninventory

import (
	"database/sql"
	"github.com/rbobillo/OnDiraitDeLaMagie/azkaban/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/azkaban/internal"
	"fmt"
)

// CreateWizards inserts a new Wizard into magicinventory
func CreatePrisoners(p dao.Prisoner, db *sql.DB) (err error) {
	populateQuery :=
		`insert into prisoners (id, magic_id)
                     values ($1, $2)`

	_, err = db.Exec(populateQuery, p.ID, p.MagicID)

	if err != nil {
		internal.Warn(fmt.Sprintf("cannot create wizard: %v , %s ", p, err))
		return err
	}

	internal.Debug(fmt.Sprintf("created wizard: %v", p))
	return nil
}

// InitAzkabanInventory function sets up the azkabaninventory db
func InitAzkabanInventory(psqlInfo string) (*sql.DB, error) {
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		internal.Warn("error while opening db connection")
		return db, err
	}
	initQuery :=
		`create table if not exists prisoners (
            id         uuid        not null primary key,
            magic_id   uuid        not null primary key,
            arrested   TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
         ); alter table prisoners owner to magic;`

	_, err = db.Query(initQuery)

	if err != nil {
		internal.Error(err.Error())
		return db, err
	}

	internal.Debug("prisoners table created")

	return db, err
}