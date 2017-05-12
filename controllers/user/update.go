package user

import (
	"net/http"

	"github.com/emurmotol/nmsrs/helpers/img"
	"github.com/emurmotol/nmsrs/helpers/lang"
	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/emurmotol/nmsrs/helpers/tpl"
	"github.com/emurmotol/nmsrs/helpers/vald"
	"github.com/emurmotol/nmsrs/models/user"
	"github.com/gorilla/mux"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	usr, err := user.FindByID(mux.Vars(r)["id"])

	if err != nil {
		panic(err)
	}
	data := map[string]interface{}{
		"Title": "Edit user",
		"User":  usr,
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "menu", "user.edit", data, funcMap)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(0); err != nil {
		panic(err)
	}
	photoFieldName := "photo"
	file, _, err := r.FormFile(photoFieldName)
	newFileInstance, handler, _ := r.FormFile(photoFieldName) // TODO: Duplicate instance of form file

	if err != http.ErrMissingFile {
		panic(err)
	}
	delete(r.PostForm, photoFieldName)
	var profile user.Profile

	if err := decoder.Decode(&profile, r.PostForm); err != nil {
		panic(err)
	}
	errs := vald.StructHasError(profile)

	id := mux.Vars(r)["id"]
	sameAsOld, err := user.CheckEmailIfSameAsOld(id, profile.Email)

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
		})
		return
	} // TODO: Validate on client

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
				panic(err)
			}
		} else {
			if err := user.SetPhoto(file, id); err != nil {
				panic(err)
			}
		}
	}

	if len(errs) != 0 {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   errs,
		})
		return
	}

	if err := user.UpdateProfile(id, profile); err != nil {
		panic(err)
	}
	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data: map[string]string{
			"message": lang.En["UserSuccessUpdate"],
		},
	})
	return
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	var resetPassword user.ResetPassword

	if err := decoder.Decode(&resetPassword, r.PostForm); err != nil {
		panic(err)
	}
	errs := vald.StructHasError(resetPassword)

	if len(errs) != 0 {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
		})
		return
	}
	id := mux.Vars(r)["id"]

	if err := user.UpdatePassword(id, resetPassword); err != nil {
		panic(err)
	}
	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data: map[string]string{
			"message": lang.En["PasswordSuccessUpdate"],
		},
	})
	return
}
