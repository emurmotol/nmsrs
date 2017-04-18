package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/zneyrl/nmsrs-lookup/helpers/flash"
	"github.com/zneyrl/nmsrs-lookup/helpers/res"
	"github.com/zneyrl/nmsrs-lookup/helpers/str"
	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
	"github.com/zneyrl/nmsrs-lookup/helpers/trans"
	"github.com/zneyrl/nmsrs-lookup/models/user"
)

var (
	decoder = schema.NewDecoder()
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
	var usr user.User

	if err := decoder.Decode(&usr, r.PostForm); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	errs := trans.StructHasError(usr)

	if len(errs) != 0 {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
			Errors: errs,
		})
		return
	}

	if err := usr.CheckEmailIfTaken(); err != nil {
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

func Edit(w http.ResponseWriter, r *http.Request) {
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
		"Title": "Edit User",
		"User":  usr,
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
	var profile user.Profile

	if err := decoder.Decode(&profile, r.PostForm); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	errs := trans.StructHasError(profile)

	if len(errs) != 0 {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
			Errors: errs,
		})
		return
	}
	v := mux.Vars(r)
	id := v["id"]
	var usr user.User
	usr.Name = profile.Name
	usr.Email = profile.Email

	if err := usr.CheckEmailIfSameAsOld(id); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	src := user.Src{
		"name":  profile.Name,
		"email": profile.Email,
	}

	if err := user.Update(id, src); err != nil {
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
			"redirect": "/users/" + id + "/edit" + r.URL.Fragment,
		},
		Errors: "",
	})
	return
}

func Destroy(w http.ResponseWriter, r *http.Request) {
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

	if err := usr.Delete(); err != nil {
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

func DestroyMany(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	var ids []string

	for _, values := range r.Form {
		for _, value := range values {
			ids = append(ids, value)
		}
	}

	if err := user.DeleteMany(ids); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}

	if err := flash.Set(r, w, "Users deleted"); err != nil {
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

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	var resetPassword user.ResetPassword

	if err := decoder.Decode(&resetPassword, r.PostForm); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	errs := trans.StructHasError(resetPassword)

	if len(errs) != 0 {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
			Errors: errs,
		})
		return
	}
	v := mux.Vars(r)
	id := v["id"]
	var usr user.User
	usr.Password = str.Bcrypt(resetPassword.Password)
	usr.ConfirmPassword = resetPassword.ConfirmPassword

	if err := usr.Update(id); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}

	if err := flash.Set(r, w, "Password updated"); err != nil {
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
			"redirect": "/users/" + id + "/edit" + r.URL.Fragment,
		},
		Errors: "",
	})
	return
}
