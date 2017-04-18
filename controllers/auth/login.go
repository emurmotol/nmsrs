package auth

import (
	"fmt"
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/env"
	"github.com/zneyrl/nmsrs-lookup/helpers/res"
	"github.com/zneyrl/nmsrs-lookup/helpers/str"
	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
	"github.com/zneyrl/nmsrs-lookup/helpers/trans"
	"github.com/zneyrl/nmsrs-lookup/middlewares"
	"github.com/zneyrl/nmsrs-lookup/models/user"
)

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Login",
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "auth", "auth.login", data, funcMap)
}

func Login(w http.ResponseWriter, r *http.Request) {
	var authCredentials user.AuthCredentials

	if err := r.ParseForm(); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}

	if err := decoder.Decode(&authCredentials, r.PostForm); err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: err.Error(),
		})
		return
	}
	errs := trans.StructHasError(authCredentials)

	if len(errs) != 0 {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
			Errors: errs,
		})
		return
	}
	usr, err := user.FindByEmail(authCredentials.Email)

	if err != nil {
		res.JSON(w, res.Make{
			Status: http.StatusInternalServerError,
			Data:   "",
			Errors: map[string]interface{}{
				"email": fmt.Sprintf("Sorry, %s doesn't recognize that email", env.AppName),
			},
		})
		return
	}

	if !str.IsPasswordMatched(usr.Password, authCredentials.Password) {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
			Errors: "These credentials do not match our records",
		})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:       middlewares.TokenName,
		Value:      middlewares.GetToken(),
		Path:       "/",
		RawExpires: "0",
	})

	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data: map[string]string{
			"redirect": "/dashboard",
			"message":  "Login successful",
		},
		Errors: "",
	})
	return
}
