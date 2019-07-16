// Package api is the entry point of 'magic' service
//
package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
)

// InitMagic starts the Magic service
func InitMagic(db *sql.DB) (err error) {
	rt := mux.NewRouter().StrictSlash(true) // handle trailing slash on each route

	type W = http.ResponseWriter
	type R = http.Request

	rt.Methods("GET").Path("/").HandlerFunc(func(w W, r *R) { err = Index(&w, r) })

	rt.Methods("GET").Path("/wizards/").HandlerFunc(func(w W, r *R) { err = GetWizards(&w, db) })
	rt.Methods("GET").Path("/wizards/{id}/").HandlerFunc(func(w W, r *R) { err = GetWizardsByID(&w, r, db) })

	rt.Methods("POST").Path("/spawn/").HandlerFunc(func(w W, r *R) { err = SpawnWizard(&w, r, db) })

	http.Handle("/", rt)

	return err
}
