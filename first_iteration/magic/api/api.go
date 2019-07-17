// Package api is the entry point of 'magic' service
//
package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
	"path"
	"runtime"
)

// InitMagic starts the Magic service
// The exposed API is documented via Swagger (https://swagger.io/docs/specification/about/)
// The swagger-ui is handled with the folder 'swaggerui'
// it contains official swagger-ui 'dist' folder components, and this API swagger.yaml
func InitMagic(db *sql.DB) (err error) {
	rt := mux.NewRouter().StrictSlash(true) // handle trailing slash on each route

	type W = http.ResponseWriter
	type R = http.Request

	// Swagger handling
	_, mainDir, _, _ := runtime.Caller(1) // get main.go's working directory
	swaggerUiDir := path.Join(path.Dir(mainDir), "api/swaggerui/")
	rt.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(swaggerUiDir))))

	// GET actions
	rt.Methods("GET").Path("/wizards/").HandlerFunc(func(w W, r *R) { err = GetWizards(&w, db) })
	rt.Methods("GET").Path("/wizards/{id}/").HandlerFunc(func(w W, r *R) { err = GetWizardsByID(&w, r, db) })

	// POST actions
	rt.Methods("POST").Path("/wizards/spawn/").HandlerFunc(func(w W, r *R) { err = SpawnWizard(&w, r, db) })

	http.Handle("/", rt)

	return err
}
