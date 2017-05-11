package auth

import (
	"net/http"

	"github.com/emurmotol/nmsrs/env"
	"github.com/emurmotol/nmsrs/helpers/lang"
	"github.com/emurmotol/nmsrs/helpers/res"
	"github.com/emurmotol/nmsrs/helpers/str"
	"github.com/emurmotol/nmsrs/helpers/tpl"
	"github.com/emurmotol/nmsrs/middlewares"
	"github.com/emurmotol/nmsrs/models/user"
)

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	if middlewares.GetAuthID(w, r) != "" {
		http.Redirect(w, r, env.URL("/"), http.StatusFound)
	} // TODO: Temporary

	data := map[string]interface{}{
		"Title": "Login",
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "auth", "auth.login", data, funcMap)
}

func Login(w http.ResponseWriter, r *http.Request) {
	if middlewares.GetAuthID(w, r) != "" {
		return
	} // TODO: Temporary

	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	usr, err := user.FindByEmail(r.PostFormValue("email"))

	if err != nil {
		panic(err)
	}

	if !str.IsPasswordMatched(usr.Password, r.PostFormValue("password")) {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data: map[string]string{
				"error": lang.En["wrong_credentials"],
			},
		})
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:       env.JWTTokenName,
		Value:      middlewares.GetToken(usr.ObjectID.Hex()),
		Path:       "/",
		RawExpires: "0",
	})
	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data: map[string]string{
			"redirect": "/",
		},
	})
	return
}
