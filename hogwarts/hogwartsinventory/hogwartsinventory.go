// Package hogwartsinventory is used to setup and manipulate
// the magic database (hogwartsinventory)
package hogwartsinventory

import (
	"database/sql"
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/dao"
	"github.com/rbobillo/OnDiraitDeLaMagie/hogwarts/internal"
)

// CreateAttack should insert in the attack table the current attack
// CreateWizards inserts a new Wizard into magicinventory
func CreateAttack(attack dao.Action, db *sql.DB) (err error) {
	attackQuery :=
		`insert into actions(id, wizard_id, action, status)
                     values ($1, $2, $3, $4);`

	_, err = db.Exec(attackQuery, attack.ID, attack.Wizard_id, attack.Action, attack.Status)

	if err != nil {
		internal.Warn(fmt.Sprintf("cannot create action: %v , %s ", attack, err))
		return err
	}

	internal.Debug(fmt.Sprintf("created action: %v", attack))
	return nil
}


// GetAllStudents should search in the hogwartsinventory and return all students
func GetAllStudents(db *sql.DB, query string) (students []dao.Student, err error) {
	rows, err := db.Query(query)

	if err == sql.ErrNoRows {
		return students, internal.ErrStudentsNotFounds
	}

	if err != nil {
		internal.Warn("cannot get all wizards")
		return students, err
	}

	for rows.Next() {
		var stud dao.Student
		err = rows.Scan(&stud.ID, &stud.MagicID, &stud.House, &stud.Status)

		if err != nil {
			internal.Warn("cannot get all wizards: error while browsing wizards")
			return students, err
		}

		students = append(students, stud)
	}

	return students, nil
}


// InitHogwartsInventory create table actions and students
// in the hogwarts database
func InitHogwartsInventory(psqlInfo string) (*sql.DB, error){
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		internal.Error("failed to establish sql connection")
		return db, err
	}

	initActionsQuery :=
		`create table if not exists actions (
	id        uuid        not null primary key,
    wizard_id uuid        not null,
    action    varchar(50) not null,
    status    varchar(50) not null
	); alter table actions owner to hogwarts;`

	_, err = db.Query(initActionsQuery)
	if err != nil {
		internal.Error("cannot create actions table")
		internal.Error(err.Error())
		return db, err
	}

	internal.Debug("actions table created")

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

	internal.Debug("students table created")

	return db, err
}