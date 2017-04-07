package auth

import (
	"net/http"
	"strings"

	"github.com/zneyrl/nmsrs-lookup/helpers/res"
	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
	"github.com/zneyrl/nmsrs-lookup/helpers/trans"
	mw "github.com/zneyrl/nmsrs-lookup/middlewares"
	"github.com/zneyrl/nmsrs-lookup/models"
)

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Login",
	}
	tmpl.RenderWithFunc(w, "auth", "auth.login", data, nil)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.AuthCredentials

	if err := r.ParseForm(); err != nil {
		res.JSON(res.Make{http.StatusInternalServerError, "", err.Error()}, w)
		return
	}

	if err := decoder.Decode(&user, r.PostForm); err != nil {
		res.JSON(res.Make{http.StatusInternalServerError, "", err.Error()}, w)
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
		"message":  "Success login",
	}, ""}, w)
	return
}
