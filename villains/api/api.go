package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitVillains() (err error) {
	rt := mux.NewRouter().StrictSlash(true) // handle trailing slash on each route

	type W = http.ResponseWriter
	type R = http.Request

	// GET actions
	rt.Methods("GET").Path("/villains/{id}").HandlerFunc(func(w W, r *R) { err = GetVillain(&w, r) })

	// POST actions
	//rt.Methods("POST").Path("/villains/spawn").HandlerFunc(func(w W, r *R) { err = SpawnVillain(&w, r) })
	rt.Methods("POST").Path("/villains/{id}/die").HandlerFunc(func(w W, r *R) { KillVillain(&w, r) })

	http.Handle("/", rt)

	return err
}