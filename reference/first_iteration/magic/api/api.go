package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" // go get -u github.com/lib/pq
)

// GetWizards function requests the Magic Inventory
// to find wizards
func GetWizards(db *sql.DB, w *http.ResponseWriter) {
	log.Println("/GetWizards")

	rows, err := db.Query("SELECT * FROM wizards")

	if err != nil {
		panic(err)
	}

	var wizards []Wizard

	for rows.Next() {
		var wz Wizard
		err = rows.Scan(&wz.ID, &wz.FirstName, &wz.LastName, &wz.Age, &wz.Category, &wz.Arrested, &wz.Dead)

		if err != nil {
			panic(err)
		}

		wizards = append(wizards, wz)
	}

	js, _ := json.Marshal(wizards)

	fmt.Fprintf(*w, string(js))
}

// Index function exposes the swagger API description
func Index(w http.ResponseWriter, r *http.Request) {
	log.Println("/Index")
	fmt.Fprintf(w, "TODO: add Swagger API documentation")
}

// InitMagic starts the Magic service
func InitMagic(db *sql.DB) {
	http.HandleFunc("/", Index)
	http.HandleFunc("/wizards", func(w http.ResponseWriter, r *http.Request) { GetWizards(db, &w) })
}

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

func main() {
	db, err := InitMagicInventory()

	if err != nil {
		panic(err)
	}

	InitMagic(db)

	defer db.Close()

	log.Fatal(http.ListenAndServe(":9090", nil))
}