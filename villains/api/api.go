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
	rt.Methods("GET").Path("/villain/{id}").HandlerFunc(func(w W, r *R) { err = GetVillain(&w, r) })

	// POST actions
	//rt.Methods("POST").Path("/villains/spawn").HandlerFunc(func(w W, r *R) { err = SpawnVillain(&w, r) })
	//rt.Methods("POST").Path("/villains/die").HandlerFunc(func(w W, r *R) { err = KillVillain(&w, r) })

	http.Handle("/", rt)

	return err
}