package user

import (
	"net/http"

	"github.com/emurmotol/nmsrs/helpers/lang"
	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/emurmotol/nmsrs/helpers/tpl"
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
	id := mux.Vars(r)["id"]

	if err := user.UpdateProfile(id, usr); err != nil {
		panic(err)
	}

	if photo != nil {
		if err := user.SetPhoto(id, photo); err != nil {
			panic(err)
		}
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
	var usr user.User

	if err := decoder.Decode(&usr, r.PostForm); err != nil {
		panic(err)
	}
	id := mux.Vars(r)["id"]

	if err := user.UpdatePassword(id, usr); err != nil {
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
