package user

import (
	"net/http"

	"github.com/emurmotol/nmsrs/helpers/tpl"
	"github.com/emurmotol/nmsrs/models/user"
	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
	users, err := user.All()

	if err != nil {
		panic(err)
	}
	data := map[string]interface{}{
		"Title": "Users",
		"Users": users,
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "menu", "user.index", data, funcMap)
}

func Show(w http.ResponseWriter, r *http.Request) {
	usr, err := user.FindByID(mux.Vars(r)["id"])

	if err != nil {
		panic(err)
	}
	data := map[string]interface{}{
		"Title": "Show user",
		"User":  usr,
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "menu", "user.show", data, funcMap)
}
