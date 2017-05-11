package user

import (
	"net/http"

	"github.com/emurmotol/nmsrs/helpers/flash"
	"github.com/emurmotol/nmsrs/helpers/img"
	"github.com/emurmotol/nmsrs/helpers/lang"
	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/emurmotol/nmsrs/helpers/tpl"
	"github.com/emurmotol/nmsrs/helpers/vald"
	"github.com/emurmotol/nmsrs/models/user"
)

func Create(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Create user",
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "menu", "user.create", data, funcMap)
}

func Store(w http.ResponseWriter, r *http.Request) {
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
	var usr user.User

	if err := decoder.Decode(&usr, r.PostForm); err != nil {
		panic(err)
	}
	errs := vald.StructHasError(usr)

	if err := user.CheckEmailIfTaken(usr.Email); err != nil {
		if _, ok := errs["email"]; !ok {
			errs["email"] = err.Error()
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
		}
	}

	if len(errs) != 0 {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
		})
		return
	}
	id, err := usr.Insert()

	if err != nil {
		panic(err)
	}

	if file != nil {
		if err := user.SetPhoto(file, id); err != nil {
			panic(err)
		}
	} // TODO: Check file != again to capture user id

	if err := flash.Set(r, w, lang.En["user_success_create"]); err != nil {
		panic(err)
	}
	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data: map[string]string{
			"redirect": "/users",
		},
	})
	return
}
