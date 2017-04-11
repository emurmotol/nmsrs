package auth

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/helpers/res"
	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
	"github.com/zneyrl/nmsrs-lookup/helpers/trans"
	"github.com/zneyrl/nmsrs-lookup/models"
)

func ShowRegisterForm(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Register",
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "auth", "auth.register", data, funcMap)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := r.ParseForm(); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}

	if err := decoder.Decode(&user, r.PostForm); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	yes, errs := trans.StructHasError(user)

	if yes {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
			Errors: errs,
		})
		return
	}
	// TODO: Do registration logic
	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data: map[string]string{
			"redirect": "/login",
			"message":  "success register",
		},
		Errors: "",
	})
	return
}
