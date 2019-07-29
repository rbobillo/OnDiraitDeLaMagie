// Package magicinventory is used to setup and manipulate
// the magic database (magicinventory)
package magicinventory

import (
	"database/sql"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/internal"
	"log"
)

//func getId(r *http.Request) string{
//	return mux.Vars(r)["id"]
//}
//type args interface {
//	getId() string
//}

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

// UpdateWizards should update a Wizard in magicinventory
func UpdateWizards(db *sql.DB, status string, value float64) (err error) {

	_, err = db.Exec(fmt.Sprintf("UPDATE wizards SET %s = %s + $1", status, status), value)

	if err != nil {
		log.Println(fmt.Sprintf("Cannot update wizards %s", status))
		return err
	}
	log.Println(fmt.Sprintf("Wizards's %s updated", status))
	return nil
}

// UpdateWizardsByID should update a single status for single Wizard in magicinventory
func UpdateWizardsByID(db *sql.DB, query string, args interface{}, wz dao.Wizard) (err error) {

	row := db.QueryRow(query, args)

	err = row.Scan(&wz.ID, &wz.FirstName, &wz.LastName, &wz.Age, &wz.Category, &wz.Arrested, &wz.Dead)

	if err != nil {
		log.Println("Cannot update wizard status")
		return err
	}
	log.Println(fmt.Sprintf("Wizards %s 's status have been updated", args))
	return nil
}

// DeleteWizardsByID should update a Wizard in magicinventory
func DeleteWizardsByID(db *sql.DB, args interface{}) (err error) {
	_, err = db.Exec("DELETE FROM wizards WHERE id = $1;", args)

	if err != nil {
		log.Println("cannot delete wizard")
		return err
	}

	log.Printf("wizards %s have been obliviated", args)
	return nil
}

// GetAllWizards should search in the magicinventory and return all wizards
func GetAllWizards(db *sql.DB, query string, wizards []dao.Wizard) (err error) {

	rows, err := db.Query(query)

	if err != nil {
		log.Println("cannot get all wizards")
		return err
	}

	for rows.Next() {
		var wz dao.Wizard
		err = rows.Scan(&wz.ID, &wz.FirstName, &wz.LastName, &wz.Age, &wz.Category, &wz.Arrested, &wz.Dead)

		if err != nil {
			log.Println("cannot get all wizards: error while browsing wizards")
			return err
		}

		wizards = append(wizards, wz)
	}
	return nil
}

// GetWizardsByID should search a wizard by ID in the magicinventory and return it
func GetWizardsByID(db *sql.DB, query string, args interface{}, wz dao.Wizard) (err error) {
	row := db.QueryRow(query, args)
	err = row.Scan(&wz.ID, &wz.FirstName, &wz.LastName, &wz.Age, &wz.Category, &wz.Arrested, &wz.Dead)

	if err != nil {
		log.Println("cannot get wizards %s", args)
		return err
	}
	log.Println("wizard %s have been found", args)
	return nil
}

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
