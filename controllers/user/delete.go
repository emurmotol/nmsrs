package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zneyrl/nmsrs-lookup/helpers/flash"
	"github.com/zneyrl/nmsrs-lookup/helpers/res"
	"github.com/zneyrl/nmsrs-lookup/models/user"
)

func Destroy(w http.ResponseWriter, r *http.Request) {
	v := mux.Vars(r)
	usr, err := user.Find(v["id"])

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: str.UpperCaseFirstChar(err.Error()),
		})
		return
	}

	if err := usr.Delete(); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: str.UpperCaseFirstChar(err.Error()),
		})
		return
	}

	if err := flash.Set(r, w, "User has been successfully deleted"); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: str.UpperCaseFirstChar(err.Error()),
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
			Errors: str.UpperCaseFirstChar(err.Error()),
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
			Errors: str.UpperCaseFirstChar(err.Error()),
		})
		return
	}

	if err := flash.Set(r, w, "User(s) has been successfully deleted"); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: str.UpperCaseFirstChar(err.Error()),
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
