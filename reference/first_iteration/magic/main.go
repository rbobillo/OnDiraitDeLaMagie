package main

import (
	"log"
	"net/http"

	"OnDiraitDeLaMagie/reference/first_iteration/magic/api"
	"OnDiraitDeLaMagie/reference/first_iteration/magic/magicinventory"
)

func main() {
	db, err := magicinventory.InitMagicInventory()

	if err != nil {
		panic(err)
	}

	err = api.InitMagic(db)

	defer db.Close()

	log.Fatal(http.ListenAndServe(":9090", nil))
}
