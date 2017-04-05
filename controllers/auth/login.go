package auth

import (
	"net/http"
	"strings"

	validator "gopkg.in/go-playground/validator.v9"

	mw "github.com/zneyrl/nmsrs-lookup/middlewares"
	"github.com/zneyrl/nmsrs-lookup/models"
	"github.com/zneyrl/nmsrs-lookup/shared/res"
	"github.com/zneyrl/nmsrs-lookup/shared/str"
	"github.com/zneyrl/nmsrs-lookup/shared/tmpl"
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
	err = validate.Struct(user)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			res.JSON(res.Make{http.StatusInternalServerError, "", err}, w)
			return
		}
		errs := make(map[string]string)

		for _, e := range err.(validator.ValidationErrors) {
			errs[str.LowerCaseFirstChar(e.Field())] = str.CamelCaseToSentenceCase(e.Translate(trans))
		}
		res.JSON(res.Make{http.StatusForbidden, "", errs}, w)
		return
	} // TODO: Create a package for this

	if strings.ToLower(user.Email) != "admin@example.com" || user.Password != "secret" {
		res.JSON(res.Make{http.StatusForbidden, "", "Invalid credentials"}, w)
		return
	}
	// TODO: Redirect to dashboard
	res.JSON(res.Make{http.StatusOK, map[string]string{
		"token":   mw.GetToken(),
		"message": "Success login!",
	}, ""}, w)
	return
}
