package user

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/emurmotol/nmsrs/helpers/flash"
	"github.com/emurmotol/nmsrs/helpers/lang"
	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/emurmotol/nmsrs/middlewares"
	"github.com/emurmotol/nmsrs/models/user"
)

func Destroy(w http.ResponseWriter, r *http.Request) {
	usr, err := user.FindByID(mux.Vars(r)["id"])

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	id := usr.ID.Hex()

	if err := usr.Delete(); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	} // TODO: UI crashes when an error occur here

	if middlewares.GetAuthID(w, r) == id {
		res.JSON(w, res.Make{
			Status: http.StatusOK,
			Data: map[string]string{
				"redirect": "/logout",
			},
			Errors: "",
		})
		return
	}

	if err := flash.Set(r, w, lang.En["user_success_delete"]); err != nil {
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
	} // TODO: UI crashes when an error occur here

	for _, id := range ids {
		if middlewares.GetAuthID(w, r) == id {
			res.JSON(w, res.Make{
				Status: http.StatusOK,
				Data: map[string]string{
					"redirect": "/logout",
				},
				Errors: "",
			})
			return
		}
	}

	if err := flash.Set(r, w, lang.En["users_success_delete"]); err != nil {
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
