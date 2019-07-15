package main

import (
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/api"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/magicinventory"
	"log"
	"net/http"
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
