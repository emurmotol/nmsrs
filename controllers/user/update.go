package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zneyrl/nmsrs-lookup/helpers/res"
	"github.com/zneyrl/nmsrs-lookup/helpers/str"
	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
	"github.com/zneyrl/nmsrs-lookup/helpers/trans"
	"github.com/zneyrl/nmsrs-lookup/models/user"
)

func Edit(w http.ResponseWriter, r *http.Request) {
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
		"Title": "Edit User",
		"User":  usr,
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "dashboard", "user.edit", data, funcMap)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil { // TODO: Must be multipart
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

	id := mux.Vars(r)["id"]
	sameAsOld, err := user.CheckEmailIfSameAsOld(id, profile.Email)

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}

	if !sameAsOld {
		if err := user.CheckEmailIfTaken(profile.Email); err != nil {
			if _, ok := errs["email"]; !ok {
				errs["email"] = str.UpperCaseFirstChar(err.Error())
			}
		}
	}

	if len(errs) != 0 {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
			Errors: errs,
		})
		return
	}

	if err := user.UpdateProfile(id, profile); err != nil {
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
			"message": "User has been successfully updated",
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
	id := mux.Vars(r)["id"]

	if err := user.UpdatePassword(id, resetPassword); err != nil {
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
			"message": "Password has been successfully updated",
		},
		Errors: "",
	})
	return
}
