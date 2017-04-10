package user

import (
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/zneyrl/nmsrs-lookup/helpers/res"
	"github.com/zneyrl/nmsrs-lookup/helpers/str"
	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
	"github.com/zneyrl/nmsrs-lookup/helpers/trans"
	"github.com/zneyrl/nmsrs-lookup/models"
)

var (
	decoder = schema.NewDecoder()
	usr     models.User
)

func Index(w http.ResponseWriter, r *http.Request) {
	users, err := usr.All()

	if err != nil {
		res.JSON(res.Make{http.StatusInternalServerError, "", err.Error()}, w)
		return
	}
	data := map[string]interface{}{
		"Title": "Users",
		"Users": users,
	}
	funcMap := template.FuncMap{
		"DateForHuman": str.DateForHuman,
	}
	tmpl.Render(w, r, "dashboard", "user.index", data, funcMap)
}

func Create(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Create User",
	}
	tmpl.Render(w, r, "dashboard", "user.create", data, nil)
}

func Store(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		res.JSON(res.Make{http.StatusInternalServerError, "", err.Error()}, w)
		return
	}

	if err := decoder.Decode(&usr, r.PostForm); err != nil {
		res.JSON(res.Make{http.StatusInternalServerError, "", err.Error()}, w)
		return
	}
	hasErr, errs := trans.ValidationHasError(usr)

	if hasErr {
		res.JSON(res.Make{http.StatusForbidden, "", errs}, w)
		return
	}

	if err := usr.Insert(); err != nil {
		res.JSON(res.Make{http.StatusInternalServerError, "", err.Error()}, w)
		return
	}
	res.JSON(res.Make{http.StatusOK, map[string]string{
		"redirect": "/users",
		"message":  "user created",
	}, ""}, w)
	return
}

func Show(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	user, err := usr.Find(v["id"])

	if err != nil {
		res.JSON(res.Make{http.StatusNotFound, "", err.Error()}, w)
		return
	}
	data := map[string]interface{}{
		"Title": "Show User",
		"User":  user,
	}
	tmpl.Render(w, r, "dashboard", "user.show", data, nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Edit User",
	}
	tmpl.Render(w, r, "dashboard", "user.edit", data, nil)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		res.JSON(res.Make{http.StatusInternalServerError, "", err.Error()}, w)
		return
	}

	if err := decoder.Decode(&usr, r.PostForm); err != nil {
		res.JSON(res.Make{http.StatusInternalServerError, "", err.Error()}, w)
		return
	}
	hasErr, errs := trans.ValidationHasError(usr)

	if hasErr {
		res.JSON(res.Make{http.StatusForbidden, "", errs}, w)
		return
	}
	v := mux.Vars(r)
	usr.Update(v["id"])
	res.JSON(res.Make{http.StatusOK, map[string]string{
		"redirect": "", // TODO: Redirect back
		"message":  "user updated",
	}, ""}, w)
	return
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	user, err := usr.Find(v["id"])

	if err != nil {
		res.JSON(res.Make{http.StatusNotFound, "", err.Error()}, w)
		return
	}
	user.Delete()
	res.JSON(res.Make{http.StatusOK, map[string]string{
		"redirect": "/users",
		"message":  "user deleted",
	}, ""}, w)
	return
}
