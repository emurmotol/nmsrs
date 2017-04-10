package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/zneyrl/nmsrs-lookup/helpers/res"
	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
	"github.com/zneyrl/nmsrs-lookup/helpers/trans"
	"github.com/zneyrl/nmsrs-lookup/models"
)

var (
	decoder = schema.NewDecoder()
)

func Index(w http.ResponseWriter, r *http.Request) {
	var usr models.User
	usrs, err := usr.All()

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
		"Users": usrs,
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "dashboard", "user.index", data, funcMap)
}

func Create(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Create User",
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "dashboard", "user.create", data, funcMap)
}

func Store(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	var usr models.User

	if err := decoder.Decode(&usr, r.PostForm); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	hasErr, errs := trans.ValidationHasError(usr)

	if hasErr {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
			Errors: errs,
		})
		return
	}

	if err := usr.Insert(); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data: map[string]string{
			"redirect": "/users",
			"message":  "user created",
		},
		Errors: "",
	})
	return
}

func Show(w http.ResponseWriter, r *http.Request) {
	var usr models.User
	v := mux.Vars(r)
	u, err := usr.Find(v["id"])

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusNotFound,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	data := map[string]interface{}{
		"Title": "Show User",
		"User":  u,
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "dashboard", "user.show", data, funcMap)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	var usr models.User
	v := mux.Vars(r)
	u, err := usr.Find(v["id"])

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusNotFound,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	data := map[string]interface{}{
		"Title": "Edit User",
		"User":  u,
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "dashboard", "user.edit", data, funcMap)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	var usr models.User

	if err := decoder.Decode(&usr, r.PostForm); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	hasErr, errs := trans.ValidationHasError(usr)

	if hasErr {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
			Errors: errs,
		})
		return
	}
	v := mux.Vars(r)
	usr.Update(v["id"])
	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data: map[string]string{
			"redirect": r.URL.Path,
			"message":  "user updated",
		},
		Errors: "",
	})
	return
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	var usr models.User
	v := mux.Vars(r)
	u, err := usr.Find(v["id"])

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusNotFound,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	u.Delete()
	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data: map[string]string{
			"redirect": "/users",
			"message":  "user deleted",
		},
		Errors: "",
	})
	return
}
