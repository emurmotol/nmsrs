package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zneyrl/nmsrs-lookup/helpers/res"
	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
	"github.com/zneyrl/nmsrs-lookup/models/user"
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
	tmpl.Render(w, r, "dashboard", "user.index", data, funcMap)
}

func Show(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	usr, err := user.Find(v["id"])

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
	tmpl.Render(w, r, "dashboard", "user.show", data, funcMap)
}