package auth

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/models"
	"github.com/zneyrl/nmsrs-lookup/shared/res"
	"github.com/zneyrl/nmsrs-lookup/shared/tmpl"
	"github.com/zneyrl/nmsrs-lookup/shared/trans"
)

func ShowRegisterForm(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Register",
	}
	tmpl.Render(w, "auth", "auth.register", data)
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := r.ParseForm()
	err = decoder.Decode(&user, r.PostForm)

	if err != nil {
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
		"message":  "Success register!",
	}, ""}, w)
	return
}
