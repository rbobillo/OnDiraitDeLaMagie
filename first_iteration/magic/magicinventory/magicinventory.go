// Package magicinventory is used to setup and manipulate
// the magic database (magicinventory)
package magicinventory

import (
	"database/sql"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/internal"
	"log"
	"strconv"
)

// CreateWizard inserts a new Wizard into magicinventory
func CreateWizard(w dao.Wizard, db *sql.DB) (err error) {
	populateQuery :=
		`insert into wizards (id, first_name, last_name, age, category, arrested, dead)
                values ($1, $2, $3, $4, $5, $6, $7);`

	_, err = db.Exec(populateQuery, w.ID, w.FirstName, w.LastName, w.Age, w.Category, w.Arrested, w.Dead)

	if err != nil {
		log.Println("Cannot create wizard:", w, err)
		return err
	}

	log.Println("Created wizard:", w)

	return nil
}

// UpdateWizard should update a Wizard in magicinventory
func UpdateWizard(w dao.Wizard, db *sql.DB, n float64) (err error) {
	updateQuery :=
		`update wizards set age = age + $n;`
	_, err = db.Exec(updateQuery, n);

	if err != nil {
		log.Println("Cannot increase wizard age")
		return err
	}
	log.Println("Wizards's age increased by " + strconv.Itoa(int(n)))
	return nil
}

// DeleteWizard should update a Wizard in magicinventory
func DeleteWizard(w dao.Wizard, db *sql.DB) (err error) { return nil }

// populateMagicInventory function should create random wizards
// and fill the magicinventory with them
func populateMagicInventory(db *sql.DB) error {
	body, _ := internal.GetRandomNames(10)
	wizards, err := internal.GenerateWizards(body)

	for _, w := range wizards {
		err = CreateWizard(w, db)

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
func InitMagicInventory(psqlInfo string) (*sql.DB, error) {
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

	err = populateMagicInventory(db)

	return db, err
}
