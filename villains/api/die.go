package api

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rbobillo/OnDiraitDeLaMagie/villains/internal"
	"net/http"
)

func KillVillain(w *http.ResponseWriter,r *http.Request){
	id := mux.Vars(r)["id"]

	internal.Info(fmt.Sprintf("Villains %s id dead", id))

	(*w).WriteHeader(204)
}
