package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/zneyrl/nmsrs-lookup/helpers/flash"
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
	yes, errs := trans.StructHasError(usr)

	if yes {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
			Errors: errs,
		})
		return
	}

	if err := usr.CheckEmail(); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
			Errors: map[string]interface{}{
				"email": err.Error(),
			},
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

	if err := flash.Set(r, w, "User created"); err != nil {
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
			Status: http.StatusInternalServerError,
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
			Status: http.StatusInternalServerError,
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
	// TODO: Create temp struct to update spacific fields

	if err := decoder.Decode(&usr, r.PostForm); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	yes, errs := trans.StructHasError(usr)

	if yes {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
			Errors: errs,
		})
		return
	}
	// TODO: Check if email exsist, ignore if same with old
	v := mux.Vars(r)

	if err := usr.Update(v["id"]); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}

	if err := flash.Set(r, w, "User updated"); err != nil {
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
			"redirect": "/users/" + v["id"] + "/edit" + r.URL.Fragment,
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
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}

	if err := u.Delete(); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}

	if err := flash.Set(r, w, "User deleted"); err != nil {
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
		},
		Errors: "",
	})
	return
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {}
