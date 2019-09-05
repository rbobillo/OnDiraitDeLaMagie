package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitFamilies() (err error) {
	rt := mux.NewRouter().StrictSlash(true) // handle trailing slash on each route

	type W = http.ResponseWriter
	type R = http.Request

	// GET actions
	rt.Methods("GET").Path("/families/{id}").HandlerFunc(func(w W, r *R) { err = GetFamilies(&w, r) })

	// POST actions
	//rt.Methods("POST").Path("/families/spawn").HandlerFunc(func(w W, r *R) { err = SpawnNewBorn(&w, r) })
	//rt.Methods("POST").Path("/families/{id}/die").HandlerFunc(func(w W, r *R) { KillVillain(&w, r) })

	http.Handle("/", rt)

	return err
}
