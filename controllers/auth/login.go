package auth

import (
	"net/http"
	"strings"

	"github.com/zneyrl/nmsrs-lookup/helpers/res"
	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
	"github.com/zneyrl/nmsrs-lookup/helpers/trans"
	mw "github.com/zneyrl/nmsrs-lookup/middlewares"
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
			Data:   "", Errors: err.Error(),
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

	if strings.ToLower(authCredentials.Email) != "admin@example.com" || authCredentials.Password != "secret" {
		res.JSON(w, res.Make{
			Status: http.StatusForbidden,
			Data:   "",
			Errors: "Invalid credentials",
		})
		return
	}
	// TODO: Redirect to dashboard
	res.JSON(w, res.Make{
		Status: http.StatusOK,
		Data: map[string]string{
			"redirect": "/",
			"token":    mw.GetToken(),
			"message":  "success login",
		},
		Errors: "",
	})
	return
}
