package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zneyrl/nmsrs/helpers/res"
	"github.com/zneyrl/nmsrs/helpers/tmpl"
	"github.com/zneyrl/nmsrs/models/user"
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
	tmpl.Render(w, r, "menu", "user.index", data, funcMap)
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
		"Title": "Show User",
		"User":  usr,
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "menu", "user.show", data, funcMap)
}
