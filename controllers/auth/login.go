package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/schema"
	mw "github.com/zneyrl/nmsrs-lookup/middlewares"
	"github.com/zneyrl/nmsrs-lookup/models"
	"github.com/zneyrl/nmsrs-lookup/shared/response"
	"github.com/zneyrl/nmsrs-lookup/shared/tmpl"
	"gopkg.in/go-playground/validator.v9"
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
		response.JSON(response.Status{http.StatusInternalServerError, "Error parsing form"}, w)
		return
	}

	var user models.AuthCredentials
	err = decoder.Decode(&user, r.PostForm)

	if err != nil {
		response.JSON(response.Status{http.StatusInternalServerError, "Error in request"}, w)
		return
	}

	validate = validator.New()
	err = validate.Struct(user)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Fprintln(w, err)
			return
		}

		for _, err := range err.(validator.ValidationErrors) {
			fmt.Fprintln(w, err.Field()+": "+err.Tag())
		}
		// TODO: Redirect back and display errors
		return
	}

	if strings.ToLower(user.Username) != "user" || user.Password != "pass" {
		response.JSON(response.Status{http.StatusForbidden, "Invalid credentials"}, w)
		return
	}
	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString, err := token.SignedString(mw.SignKey)

	if err != nil {
		response.JSON(response.Status{http.StatusInternalServerError, "Error while signing the token"}, w)
		mw.Fatal(err)
	}
	response.JSON(response.Status{http.StatusOK, tokenString}, w)
	return
}
