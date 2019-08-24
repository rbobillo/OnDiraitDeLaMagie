// Package hogwartsinventory is used to setup and manipulate
// the magic database (hogwartsinventory)
package hogwartsinventory

import (
	"database/sql"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/internal"
	"log"
)

func InitHogwartsInventory(psqlInfo string) (*sql.DB, error){
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		internal.Error("failed to establish sql connection")
		return db, err
	}
	log.Println(db)
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	initActionsQuery :=
		`create table if not exists actions (
	id        uuid        not null primary key,
    wizard_id uuid        not null,
    category  varchar(50) not null,
    status    varchar(50) not null
	); alter table actions owner to hogwarts;`

	_, err = db.Query(initActionsQuery)
	if err != nil {
		internal.Error("cannot create actions table")
		internal.Error(err.Error())
		return db, err
	}

	internal.Info("actions table created")

	initStudentsQuery :=
		`create table if not exists students (
	id       uuid	     not null primary key,
	magic_id uuid        not null,
	house    varchar(50) not null,
    status   varchar(50) not null
    ); alter table students owner to hogwarts;`

	_, err = db.Query(initStudentsQuery)
	if err != nil {
		internal.Error("cannot create students table")
		return db, err
	}

	internal.Debug("Students table created")

	return db, err
}