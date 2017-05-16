package user

import (
	"net/http"

	"github.com/emurmotol/nmsrs/helpers/flash"
	"github.com/emurmotol/nmsrs/helpers/lang"
	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/emurmotol/nmsrs/helpers/tpl"
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
	photo, _, err := r.FormFile(photoFieldName)

	if err != nil {
		if err != http.ErrMissingFile {
			panic(err)
		}
	}
	delete(r.PostForm, photoFieldName)
	var usr user.User

	if err := decoder.Decode(&usr, r.PostForm); err != nil {
		panic(err)
	}
	id, err := usr.Insert()

	if err != nil {
		panic(err)
	}

	if photo != nil {
		if err := user.SetPhoto(id, photo); err != nil {
			panic(err)
		}
	}

	if err := flash.Set(r, w, lang.En["UserSuccessCreate"]); err != nil {
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
