package user

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/helpers/flash"
	"github.com/zneyrl/nmsrs-lookup/helpers/res"
	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
	"github.com/zneyrl/nmsrs-lookup/helpers/trans"
	"github.com/zneyrl/nmsrs-lookup/models/user"
)

func Create(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Create User",
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "dashboard", "user.create", data, funcMap)
}

func Store(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(0); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	var usr user.User
	// TODO: Populate user struct

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

	if err := user.CheckEmailIfTaken(usr.Email); err != nil {
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

	if err := flash.Set(r, w, "User has been successfully created"); err != nil {
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
