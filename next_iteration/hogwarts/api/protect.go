package api

import (
	"database/sql"
	"log"
	"net/http"
)

// ProtectHogwarts function cancels or avoids
// villains attack on Hogwarts
func ProtectHogwarts(w *http.ResponseWriter, r *http.Request, db *sql.DB) (err error) {
	log.Println("/hogwarts/protect")
	(*w).Header().Set("Content-Type", "application/json; charset=UTF-8")

	// TODO: parse body
	// TODO: create status in hogwartsinventory ?

	println(r, db)

	return nil
}
