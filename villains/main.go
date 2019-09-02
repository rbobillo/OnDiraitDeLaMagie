package main

import (
	"github.com/rbobillo/OnDiraitDeLaMagie/villains/api"
	"github.com/rbobillo/OnDiraitDeLaMagie/villains/internal"
	"log"
	"net/http"
)


func main() {
	internal.Debug("starting villains micro-service. Waiting for orders...")
	err := api.InitVillains()

	if err != nil {
		internal.Warn(err.Error())
	}

	log.Fatal(http.ListenAndServe(":9094", nil))
}
