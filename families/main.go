package main

import (
	"github.com/rbobillo/OnDiraitDeLaMagie/families/api"
	"github.com/rbobillo/OnDiraitDeLaMagie/families/internal"
	"log"
	"net/http"
)

func main(){
	internal.Debug("starting families micro-service. Waiting for event...")
	err := api.InitFamilies()

	if err != nil {
		internal.Warn(err.Error())
	}

	log.Fatal(http.ListenAndServe(":9092", nil))
}
