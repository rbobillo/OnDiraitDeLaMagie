// magic_inventory is used to setup and manipulate
// the magic database (magic_inventory)
package magic_inventory

import (
	"database/sql"
	"fmt"
	"log"
)

// TODO: these const should not be hardcoded
const (
	host     = "magic_inventory"
	port     = 5432
	user     = "magic"
	password = "magic"
	dbname   = "magic_inventory"
)

// PopulateMagicInventory function should create random users
// and fill the magic_inventory with them
func PopulateMagicInventory(db *sql.DB) error {
	body, _ := GetRandomNames(10)
	wizards, err := GenerateWizards(body)

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

// InitMagicInventory function sets up the magic_inventory db
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
                    id         varchar(50) not null primary key,
                    first_name varchar(50) not null,
                    last_name  varchar(50) not null,
                    age        float       not null,
                    category   varchar(50) not null,
                    arrested   boolean     not null,
                    dead       boolean     not null
                );
                alter table wizards owner to magic;`

	_, err = db.Query(initQuery)

	if err != nil {
		panic(err)
	}

	log.Println("Wizards table created")

	err = PopulateMagicInventory(db)

	return db, err
}