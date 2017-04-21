package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/zneyrl/nmsrs-lookup/models/user"
)

var (
	decoder = schema.NewDecoder()
)

func Photo(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, user.GetPhoto(mux.Vars(r)["id"]))
}
