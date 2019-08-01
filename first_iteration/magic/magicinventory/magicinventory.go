// Package magicinventory is used to setup and manipulate
// the magic database (magicinventory)
package magicinventory

import (
	"database/sql"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/internal"
)

// CreateWizards inserts a new Wizard into magicinventory
func CreateWizards(w dao.Wizard, db *sql.DB) (err error) {
	populateQuery :=
		`insert into wizards (id, first_name, last_name, age, category, arrested, dead)
                     values ($1, $2, $3, $4, $5, $6, $7);`

	_, err = db.Exec(populateQuery, w.ID, w.FirstName, w.LastName, w.Age, w.Category, w.Arrested, w.Dead)

	if err != nil {
		internal.Log(fmt.Sprintf("cannot create wizard: %s", err )).Error()
		return err
	}

	internal.Log(fmt.Sprintf("Created wizard: ", w )).Debug()

	return nil
}

// DeleteWizardsByID should update a Wizard in magicinventory
func DeleteWizardsByID(db *sql.DB, id string) (err error) {
	_, err = db.Exec("DELETE FROM wizards WHERE id = $1;", id)

	if err != nil {
		internal.Log(fmt.Sprintf("cannot delete wizard %s", id)).Error()
		return err
	}

	internal.Log(fmt.Sprintf("Wizards %s have been deleted", id)).Debug()

	return nil
}

// GetAllWizards should search in the magicinventory and return all wizards
func GetAllWizards(db *sql.DB, query string) (wizards []dao.Wizard, err error) {
	rows, err := db.Query(query)

	if err == sql.ErrNoRows {
		return wizards, internal.ErrWizardsNotFounds
	}

	if err != nil {
		internal.Log(fmt.Sprintf("cannot get all wizards")).Error()
		return wizards, err
	}

	for rows.Next() {
		var wz dao.Wizard
		err = rows.Scan(&wz.ID, &wz.FirstName, &wz.LastName, &wz.Age, &wz.Category, &wz.Arrested, &wz.Dead)

		if err != nil {
			internal.Log(fmt.Sprintf("cannot get all wizards: error while browsing wizards")).Error()
			return wizards, err
		}

		wizards = append(wizards, wz)
	}

	return wizards, nil
}

// GetWizardsByID should search a wizard by ID in the magicinventory and return it
func GetWizardsByID(db *sql.DB, query string, id string) (wz dao.Wizard, err error) {
	row := db.QueryRow(query, id)
	err = row.Scan(&wz.ID, &wz.FirstName, &wz.LastName, &wz.Age, &wz.Category, &wz.Arrested, &wz.Dead)

	if err == sql.ErrNoRows{
		return wz, internal.ErrWizardsNotFounds
	}
	if err != nil {
		internal.Log(fmt.Sprintf("cannot get wizards %s", id)).Error()

		return wz, err
	}

	internal.Log(fmt.Sprintf("Wizard %s have been found", id)).Debug()

	return wz, nil
}

// InitMagicInventory function sets up the magicinventory db
// TODO: use `gORM` rather than `pq` ?
// TODO: add an event listener ? https://godoc.org/github.com/lib/pq/example/listen
func InitMagicInventory(psqlInfo string) (*sql.DB, error) {
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		internal.Log(fmt.Sprintf("error while opening db connection")).Error()
		return db, err
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

	internal.Log(fmt.Sprintf("Wizards table created")).Debug()

	err = populateMagicInventory(db)

	return db, err
}

// UpdateWizards should update a Wizard in magicinventory
func UpdateWizards(db *sql.DB, query string, args ...interface{}) (err error) {
	_, err = db.Exec(query, args...)

	if err == sql.ErrNoRows {
		return  internal.ErrWizardsNotFounds
	}

	if err != nil {
		internal.Log(fmt.Sprintf("cannot update wizards")).Error()
		return err
	}

	internal.Log(fmt.Sprintf("Wizards updated")).Debug()

	return nil
}

// UpdateWizardsByID should update a single status for single Wizard in magicinventory
func UpdateWizardsByID(db *sql.DB, id string, query string, args ...interface{}) (wz dao.Wizard, err error) {
	args = append([]interface{}{id}, args...)
	row := db.QueryRow(query, args...)

	err = row.Scan(&wz.ID, &wz.FirstName, &wz.LastName, &wz.Age, &wz.Category, &wz.Arrested, &wz.Dead)

	if err == sql.ErrNoRows {
		return wz, internal.ErrWizardsNotFounds
	}

	if err != nil {
		internal.Log(fmt.Sprintf("cannot update wizard status")).Error()
		return wz, err
	}

	internal.Log(fmt.Sprintf("Wizard %s's status has been updated", id)).Debug()

	return wz, err
}

// populateMagicInventory function should create random wizards
// and fill the magicinventory with them
func populateMagicInventory(db *sql.DB) error {
	body, _ := internal.GetRandomNames(10)
	wizards, err := internal.GenerateWizards(body)

	for _, w := range wizards {
		err = CreateWizards(w, db)

		if err != nil {
			internal.Log(fmt.Sprintf("cannot populate magic inventory")).Error()
			return err
		}
	}

	internal.Log(fmt.Sprintf("Wizards table populated")).Debug()
	return err
}
