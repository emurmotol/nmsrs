package auth

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/models"
	"github.com/zneyrl/nmsrs-lookup/shared/res"
	"github.com/zneyrl/nmsrs-lookup/shared/str"
	"github.com/zneyrl/nmsrs-lookup/shared/tmpl"
	validator "gopkg.in/go-playground/validator.v9"
)

func ShowRegisterForm(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Register",
	}
	tmpl.Render(w, "auth", "auth.register", data)
}

func Register(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		res.JSON(res.Make{http.StatusInternalServerError, "", "Error parsing form"}, w)
		return
	}
	var user models.User
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
	// TODO: Do registration logic
	res.JSON(res.Make{http.StatusOK, map[string]string{
		"message": "Success register!",
	}, ""}, w)
	return
}
