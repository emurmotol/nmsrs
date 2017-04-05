package auth

import (
	"net/http"
	"strings"
	"time"

	validator "gopkg.in/go-playground/validator.v9"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/schema"
	mw "github.com/zneyrl/nmsrs-lookup/middlewares"
	"github.com/zneyrl/nmsrs-lookup/models"
	"github.com/zneyrl/nmsrs-lookup/shared/response"
	"github.com/zneyrl/nmsrs-lookup/shared/tmpl"
)

var decoder = schema.NewDecoder()
var validate *validator.Validate

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Login",
	}
	tmpl.Render(w, "auth", "auth.login", data)
}

func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		response.JSON(response.Make{http.StatusInternalServerError, "", "Error parsing form"}, w)
		return
	}

	var user models.AuthCredentials
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
			// TODO: Fix validation message for other tags
			// TODO: Convert camel case to snake case
			errs[strings.ToLower(err.Field())] = err.Field() + " is " + err.Tag()
		}
		// TODO: Redirect back and display errors
		response.JSON(response.Make{http.StatusForbidden, "", errs}, w)
		return
	} // TODO: Create a package for this

	if strings.ToLower(user.Username) != "user" || user.Password != "pass" {
		response.JSON(response.Make{http.StatusForbidden, "", "Invalid credentials"}, w)
		return
	}
	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString, err := token.SignedString(mw.SignKey)

	if err != nil {
		response.JSON(response.Make{http.StatusInternalServerError, "", "Error while signing the token"}, w)
		mw.Fatal(err)
	}
	response.JSON(response.Make{http.StatusOK, map[string]string{
		"token": tokenString,
	}, ""}, w)
	return
}
