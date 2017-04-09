package user

import (
	"net/http"

	"github.com/gorilla/schema"
	"github.com/zneyrl/nmsrs-lookup/helpers/res"
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
		"R":     r,
		"Users": users,
	}
	tmpl.RenderWithFunc(w, "dashboard", "user.index", data, nil)
}

func Create(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Create User",
		"R":     r,
	}
	tmpl.RenderWithFunc(w, "dashboard", "user.create", data, nil)
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
		"message":  "User created",
	}, ""}, w)
	return
}

func Show(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Show User",
		"R":     r,
	}
	tmpl.RenderWithFunc(w, "dashboard", "user.show", data, nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Edit User",
		"R":     r,
	}
	tmpl.RenderWithFunc(w, "dashboard", "user.edit", data, nil)
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
	// TODO: models.User.Update()
	res.JSON(res.Make{http.StatusOK, map[string]string{
		"redirect": "", // TODO: Redirect back?
		"message":  "User updated",
	}, ""}, w)
	return
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	// TODO: models.User.Delete()
}
