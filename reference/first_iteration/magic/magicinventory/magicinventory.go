// Package magicinventory is used to setup and manipulate
// the magic database (magicinventory)
package magicinventory

import (
	"database/sql"
	"fmt"
	"log"

	"OnDiraitDeLaMagie/reference/first_iteration/magic/internal"
)

// TODO: these const should not be hardcoded
const (
	host     = "localhost" // "magicinventory" // or localhost (if you run the service outside of Docker)
	port     = 5432
	user     = "magic"
	password = "magic"
	dbname   = "magicinventory"
)

// PopulateMagicInventory function should create random users
// and fill the magicinventory with them
func PopulateMagicInventory(db *sql.DB) error {
	body, _ := internal.GetRandomNames(10)
	wizards, err := internal.GenerateWizards(body)

	populateQuery :=
		`insert into wizards (id, first_name, last_name, age, category, arrested, dead)
                values ($1, $2, $3, $4, $5, $6, $7);`

	for _, w := range wizards {
		_, err = db.Exec(populateQuery, w.ID, w.FirstName, w.LastName, w.Age, w.Category, w.Arrested, w.Dead)

		if err != nil {
			return err
		}
	}

	log.Println("Wizards table populated")

	return err
}

// InitMagicInventory function sets up the magicinventory db
// TODO: use `gORM` rather than `pq` ?
// TODO: add an event listener ? https://godoc.org/github.com/lib/pq/example/listen
func InitMagicInventory() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s "+
		"dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	initQuery :=
		`create table if not exists wizards (
            id         uuid        not null primary key,
            first_name varchar(50) not null,
            last_name  varchar(50) not null,
            age        float       not null,
        	category   varchar(50) not null,
            arrested   boolean     not null,
            dead       boolean     not null
         ); alter table wizards owner to magic;`

	_, err = db.Query(initQuery)

	if err != nil {
		panic(err)
	}

	log.Println("Wizards table created")

	err = PopulateMagicInventory(db)

	return db, err
}
