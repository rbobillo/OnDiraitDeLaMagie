// Package api is the entry point of 'magic' service
//
package api

import (
	"github.com/rbobillo/OnDiraitDeLaMagie/first_iteration/magic/internal"
	"net/http"
	"fmt"
	_ "github.com/lib/pq" // go get -u github.com/lib/pq
)

// Index function exposes the swagger API description
func Index(w *http.ResponseWriter, r *http.Request) (err error) {
	internal.Log(fmt.Sprintf("/ redirect to /swaggerui/")).Debug()

	http.Redirect(*w, r, "/swaggerui/", http.StatusFound)

	return err
}
