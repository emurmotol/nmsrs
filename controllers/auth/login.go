package auth

import (
	"net/http"
	"strings"

	mw "github.com/zneyrl/nmsrs-lookup/middlewares"
	"github.com/zneyrl/nmsrs-lookup/models"
	"github.com/zneyrl/nmsrs-lookup/shared/res"
	"github.com/zneyrl/nmsrs-lookup/shared/tmpl"
	"github.com/zneyrl/nmsrs-lookup/shared/trans"
)

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Login",
	}
	tmpl.Render(w, "auth", "auth.login", data)
}

func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		res.JSON(res.Make{http.StatusInternalServerError, "", "Error parsing form"}, w)
		return
	}
	var user models.AuthCredentials
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

	if strings.ToLower(user.Email) != "admin@example.com" || user.Password != "secret" {
		res.JSON(res.Make{http.StatusForbidden, "", "Invalid credentials"}, w)
		return
	}
	// TODO: Redirect to dashboard
	res.JSON(res.Make{http.StatusOK, map[string]string{
		"redirect": "/",
		"token":    mw.GetToken(),
		"message":  "Success login!",
	}, ""}, w)
	return
}
