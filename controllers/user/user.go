package user

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/zneyrl/nmsrs-lookup/models/user"
)

var (
	decoder = schema.NewDecoder()
)

func Photo(w http.ResponseWriter, r *http.Request) {
	file := path.Join(user.ContentDir, mux.Vars(r)["id"], "photo", "default.jpg") // TODO: File extension must be dynamic
	http.ServeFile(w, r, file)
}
