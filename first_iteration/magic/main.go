package main

import (
	"fmt"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/api"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/internal"
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/magicinventory"
	"log"
	"net/http"
)

func main() {
	hostname := internal.GetEnvOrElse("PG_HOST", "localhost")
	portaddr := internal.GetEnvOrElse("PG_PORT", "5432")
	username := internal.GetEnvOrElse("POSTGRES_USER", "magic")
	password := internal.GetEnvOrElse("POSTGRES_PASSWORD", "magic")
	database := internal.GetEnvOrElse("POSTGRES_DB", "magicinventory")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		hostname, portaddr, username, password, database)

	db, err := magicinventory.InitMagicInventory(psqlInfo)

	if err != nil {
		internal.Error(err.Error())
	}

	err = api.InitMagic(db)

	defer db.Close()

	log.Fatal(http.ListenAndServe(":9090", nil))
}
