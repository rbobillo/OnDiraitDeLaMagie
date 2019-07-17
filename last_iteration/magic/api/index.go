// Package api is the entry point of 'magic' service
//
package api

import (
	"log"
	"net/http"

	_ "github.com/lib/pq" // go get -u github.com/lib/pq
)

// Index function exposes the swagger API description
func Index(w *http.ResponseWriter, r *http.Request) (err error) {
	log.Println("/ redirect to /swaggerui/")
	http.Redirect(*w, r, "/swaggerui/", http.StatusFound)

	return err
}
