package auth

import (
	"net/http"

	validator "gopkg.in/go-playground/validator.v9"

	"github.com/zneyrl/nmsrs-lookup/models"
	"github.com/zneyrl/nmsrs-lookup/shared/response"
	"github.com/zneyrl/nmsrs-lookup/shared/str"
	"github.com/zneyrl/nmsrs-lookup/shared/tmpl"
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
		response.JSON(response.Make{http.StatusInternalServerError, "", "Error parsing form"}, w)
		return
	}
	var user models.User
	err = decoder.Decode(&user, r.PostForm)

	if err != nil {
		response.JSON(response.Make{http.StatusInternalServerError, "", "Error in request"}, w)
		return
	}
	validate = validator.New()
	err = validate.Struct(user)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			response.JSON(response.Make{http.StatusInternalServerError, "", err}, w)
			return
		}
		errs := make(map[string]string)

		for _, err := range err.(validator.ValidationErrors) {
			// TODO: Make translation of messages
			// TODO: Convert title case to sentence case
			errs[str.LowerCaseFirstChar(err.Field())] = err.Field() + " is " + err.Tag()
		}
		// TODO: Redirect back and display errors
		response.JSON(response.Make{http.StatusForbidden, "", errs}, w)
		return
	} // TODO: Create a package for this

	if r.PostFormValue("password") != r.PostFormValue("confirmPassword") {
		response.JSON(response.Make{http.StatusForbidden, "", map[string]string{
			"confirmPassword": "Confirm password did not match",
		}}, w)
		return
	}
	// TODO: Do registration logic
	response.JSON(response.Make{http.StatusOK, map[string]string{
		"message": "Success register!",
	}, ""}, w)
	return
}
