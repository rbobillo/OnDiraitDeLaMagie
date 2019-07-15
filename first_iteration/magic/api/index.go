// Package api is the entry point of 'magic' service
//
package api

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq" // go get -u github.com/lib/pq
)

// Index function exposes the swagger API description
func Index(w *http.ResponseWriter, r *http.Request) (err error) {
	log.Println("/")
	_, err = fmt.Fprintf(*w, "TODO: add Swagger API documentation")

	return err
}