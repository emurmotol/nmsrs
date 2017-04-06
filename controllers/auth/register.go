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
	tmpl.RenderWithFunc(w, "auth", "auth.register", data, nil)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := r.ParseForm(); err != nil {
		res.JSON(res.Make{http.StatusInternalServerError, "", "Error parsing form"}, w)
		return
	}

	if err := decoder.Decode(&user, r.PostForm); err != nil {
		res.JSON(res.Make{http.StatusInternalServerError, "", "Error in request"}, w)
		return
	}
	hasErr, errs := trans.ValidationHasError(user)

	if hasErr {
		res.JSON(res.Make{http.StatusForbidden, "", errs}, w)
		return
	}
	// TODO: Do registration logic
	res.JSON(res.Make{http.StatusOK, map[string]string{
		"redirect": "/login",
		"message":  "Success register",
	}, ""}, w)
	return
}
