// Package api is the entry point of 'magic' service
//
package api

import (
	_ "github.com/lib/pq" // go get -u github.com/lib/pq
	"github.com/rbobillo/OnDiraitDeLaMagie/magic/internal"
	"net/http"
)

// Index function exposes the swagger API description
func Index(w *http.ResponseWriter, r *http.Request) (err error) {
	internal.Debug("/ redirect to /swaggerui/")

	http.Redirect(*w, r, "/swaggerui/", http.StatusFound)

	return err
}
