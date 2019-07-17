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

// function ServeSwaggerUI allows to display
// API documentation on localhost:9090/
// It find current working dir (cwd)
// And joins it to swagger-ui dir relative path
// Then it loads this dir (and trims it from serve path)
func ServeSwaggerUI(rt *mux.Router) {
	_, cwd, _, _ := runtime.Caller(1)

	ui := path.Join(path.Dir(cwd), "/swaggerui/")

	rt.PathPrefix("/swaggerui/").Handler(
		http.StripPrefix("/swaggerui/", http.FileServer(http.Dir(ui))))
}

// InitMagic starts the Magic service
// The exposed API is documented via Swagger (https://swagger.io/docs/specification/about/)
// The swagger-ui is handled with the folder 'swaggerui'
// it contains official swagger-ui 'dist' folder components, and this API swagger.yaml
func InitMagic(db *sql.DB) (err error) {
	rt := mux.NewRouter().StrictSlash(true) // handle trailing slash on each route

	type W = http.ResponseWriter
	type R = http.Request

	// Swagger handling
	ServeSwaggerUI(rt)

	// GET actions
	rt.Methods("GET").Path("/").HandlerFunc(func(w W, r *R) { err = Index(&w, r) })
	rt.Methods("GET").Path("/wizards/").HandlerFunc(func(w W, r *R) { err = GetWizards(&w, db) })
	rt.Methods("GET").Path("/wizards/{id}/").HandlerFunc(func(w W, r *R) { err = GetWizardsByID(&w, r, db) })

	// POST actions
	rt.Methods("POST").Path("/wizards/spawn/").HandlerFunc(func(w W, r *R) { err = SpawnWizard(&w, r, db) })

	http.Handle("/", rt)

	return err
}
