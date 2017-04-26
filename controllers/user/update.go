package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zneyrl/nmsrs/helpers/img"
	"github.com/zneyrl/nmsrs/helpers/lang"
	"github.com/zneyrl/nmsrs/helpers/res"
	"github.com/zneyrl/nmsrs/helpers/tpl"
	"github.com/zneyrl/nmsrs/helpers/vald"
	"github.com/zneyrl/nmsrs/models/user"
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
	tpl.Render(w, r, "menu", "user.edit", data, funcMap)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(0); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	photoFieldName := "photo"
	file, _, err := r.FormFile(photoFieldName)
	newFileInstance, handler, _ := r.FormFile(photoFieldName) // TODO: Duplicate instance of form file

	if err != nil {
		if err != http.ErrMissingFile {
			res.JSON(w, res.Make{
				Status: http.StatusInternalServerError,
				Data:   "",
				Errors: err.Error(),
			})
			return
		}
	}
	delete(r.PostForm, photoFieldName)
	var profile user.Profile

	if err := decoder.Decode(&profile, r.PostForm); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	errs := vald.StructHasError(profile)

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
				errs["email"] = err.Error()
			}
		}
	}

	if file != nil {
		if err := img.Validate(newFileInstance, handler); err != nil {
			if err == img.ErrImageNotValid || err == img.ErrImageTooLarge { // TODO: Add new custom err here
				if _, ok := errs[photoFieldName]; !ok {
					errs[photoFieldName] = err.Error()
				}
			} else {
				res.JSON(w, res.Make{
					Status: http.StatusInternalServerError,
					Data:   "",
					Errors: err.Error(),
				})
				return
			}
		} else {
			if err := user.SetPhoto(file, id); err != nil {
				res.JSON(w, res.Make{
					Status: http.StatusInternalServerError,
					Data:   "",
					Errors: err.Error(),
				})
				return
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
			"message": lang.En["user_success_update"],
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
	errs := vald.StructHasError(resetPassword)

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
			"message": lang.En["password_success_update"],
		},
		Errors: "",
	})
	return
}
