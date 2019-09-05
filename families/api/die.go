package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rbobillo/OnDiraitDeLaMagie/families/internal"
	"net/http"
)

func KillVillain(w *http.ResponseWriter,r *http.Request){
	id := mux.Vars(r)["id"]

	internal.Info(fmt.Sprintf("Family wizard %s id dead", id))

	(*w).WriteHeader(204)
}
