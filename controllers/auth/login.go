package auth

import (
	"net/http"

	"github.com/emurmotol/nmsrs/env"
	"github.com/emurmotol/nmsrs/helpers/lang"
	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/emurmotol/nmsrs/helpers/str"
	"github.com/emurmotol/nmsrs/helpers/tpl"
	"github.com/emurmotol/nmsrs/helpers/vald"
	"github.com/emurmotol/nmsrs/middlewares"
	"github.com/emurmotol/nmsrs/models/user"
)

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	if middlewares.GetAuthID(r) != "" {
		http.Redirect(w, r, env.URL("/"), http.StatusFound)
	} // TODO: Temporary

	data := map[string]interface{}{
		"Title": "Login",
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "auth", "auth.login", data, funcMap)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if middlewares.GetAuthID(r) != "" {
		return
	} // TODO: Temporary

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
	errs := vald.StructHasError(authCredentials)
	usr, err := user.FindByEmail(authCredentials.Email)

	if err != nil {
		if _, ok := errs["email"]; !ok {
			errs["email"] = lang.En["email_not_recognized"]
		}
	}

	if len(errs) != 0 {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
			Errors: errs,
		})
		return
	}

	if !str.IsPasswordMatched(usr.Password, authCredentials.Password) {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
			Errors: lang.En["wrong_credentials"],
		})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:       env.JWTTokenName,
		Value:      middlewares.GetToken(usr.ID.Hex()),
		Path:       "/",
		RawExpires: "0",
	})
	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data: map[string]string{
			"redirect": "/",
			"message":  lang.En["user_authenticated"],
		},
		Errors: "",
	})
	return
}
