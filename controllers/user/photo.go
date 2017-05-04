package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/emurmotol/nmsrs/models/user"
)

func Photo(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, user.GetPhoto(mux.Vars(r)["id"]))
}
