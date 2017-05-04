package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/emurmotol/nmsrs/helpers/tpl"
	"github.com/emurmotol/nmsrs/models/user"
)

func Index(w http.ResponseWriter, r *http.Request) {
	users, err := user.All()

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	data := map[string]interface{}{
		"Title": "Users",
		"Users": users,
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "menu", "user.index", data, funcMap)
}

func Show(w http.ResponseWriter, r *http.Request) {
	usr, err := user.Find(mux.Vars(r)["id"])

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	data := map[string]interface{}{
		"Title": "Show user",
		"User":  usr,
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "menu", "user.show", data, funcMap)
}
