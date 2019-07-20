// Package hogwartsinventory is used to setup and manipulate
// the magic database (hogwartsinventory)
package hogwartsinventory

import (
	"database/sql"
	//"github.com/rbobillo/OnDiraitDeLaMagie/next_iteration/hogwarts/dao"
	//"github.com/rbobillo/OnDiraitDeLaMagie/next_iteration/hogwarts/internal"
	"log"
)

// InitHogwartsInventory function sets up the hogwartsinventory db
func InitHogwartsInventory(psqlInfo string) (*sql.DB, error) {
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	initQuery :=
		`create table if not exists students (
            id         uuid        not null primary key,
            magic_id   uuid        not null,
            house      varchar(50) not null,
            status     varchar(50) not null
         ); alter table students owner to magic;`

	_, err = db.Query(initQuery)

	if err != nil {
		panic(err)
	}

	log.Println("Students table created")

	// TODO: create & populate other tables ?

	return db, err
}
