package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitGuest()(err error) {
	rt := mux.NewRouter().StrictSlash(true) // handle trailing slash on each route

	type W = http.ResponseWriter
	type R = http.Request

	// GET actions
	rt.Methods("GET").Path("/guest/{id}").HandlerFunc(func(w W, r *R) { err = GetGuestByID(&w, r) })

	//// POST actions
	rt.Methods("POST").Path("/guest/spawn").HandlerFunc(func(w W, r *R) { err = StartSlotEvent(&w, r) })
	//rt.Methods("POST").Path("/guest/{id}/die").HandlerFunc(func(w W, r *R) { KillVillain(&w, r) })

	http.Handle("/", rt)

	return err
}
