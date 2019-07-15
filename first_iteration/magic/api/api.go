// Package api is the entry point of 'magic' service
//
package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"

	_ "github.com/lib/pq" // go get -u github.com/lib/pq
)

// InitMagic starts the Magic service
func InitMagic(db *sql.DB) (err error) {
	rt := mux.NewRouter().StrictSlash(true) // handle trailing slash on each route

	type W = http.ResponseWriter
	type R = http.Request

	rt.Methods("GET").Path("/").HandlerFunc(func(w W, r *R) { err = Index(&w, r) })

	rt.Methods("GET").Path("/wizards/").HandlerFunc(func(w W, r *R) { err = GetWizards(&w, db) })
	rt.Methods("GET").Path("/wizards/{id}/").HandlerFunc(func(w W, r *R) { err = GetWizardsById(&w, r, db) })

	http.Handle("/", rt)

	return err
}
