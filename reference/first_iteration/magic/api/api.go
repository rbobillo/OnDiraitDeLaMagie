package main

import (
	//"database/sql"
	"fmt"
	"log"
	"net/http"

	// _ "github.com/lib/pq" // go get -u github.com/lib/pq
)

// GetWizards function requests the Magic Inventory
// to find wizards
func GetWizards(w *http.ResponseWriter, r *http.Request) error {
	fmt.Fprintf(*w, "Getting Wizards")
	return nil
}

// InitMagic starts the Magic service
func InitMagic() error {
	fmt.Println("Init Magic service")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "TODO: add Swagger API documentation")
	})


	http.HandleFunc("/wizards", func(w http.ResponseWriter, r *http.Request){
		GetWizards(&w, r) //fmt.Fprintf(w, "no wizard yet")
	})

	return nil
}

/*
const (
	host     = "localhost"
	port     = 5432
	user     = "magic"
	password = "magic"
	dbname   = "magic_inventory"
)

func InitMagicInventory() error {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s " +
	"dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	return err
}
*/

func main() {
	var err error = nil

	/*err = InitMagicInventory()

	if err != nil {
		panic(err)
	}*/

	err = InitMagic()

	if err != nil {
		panic(err)
	}

	log.Fatal(http.ListenAndServe(":9090", nil))
}
