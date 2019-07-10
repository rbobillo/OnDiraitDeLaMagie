package main

import (
	"log"
	"net/http"

	"github.com/rbobillo/OnDiraitDeLaMagie/reference/first_iteration/magic/api"
	"github.com/rbobillo/OnDiraitDeLaMagie/reference/first_iteration/magic/magic_inventory"
)

func main() {
	db, err := InitMagicInventory()

	if err != nil {
		panic(err)
	}

	InitMagic(db)

	defer db.Close()

	log.Fatal(http.ListenAndServe(":9090", nil))
}