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
<<<<<<< HEAD
var validate *validator.Validate
=======
>>>>>>> 8e4ec4c41d89c9406d3c186dddc3e1129455dab6

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Login",
	}
	tmpl.Render(w, "auth", "auth.login", data)
}

func Login(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD
	err := r.ParseForm()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error parsing form")
		return
	}

	var user models.AuthCredentials
	err = decoder.Decode(&user, r.PostForm)
=======
	r.ParseForm() // TODO: Must handle error
	var user models.UserCredentials
	err := decoder.Decode(&user, r.PostForm)
>>>>>>> 8e4ec4c41d89c9406d3c186dddc3e1129455dab6

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error in request")
		return
	}
<<<<<<< HEAD

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
		// TODO: Redirect back
		return
	}

	if strings.ToLower(user.Username) != "user" || user.Password != "pass" {
=======

	if strings.ToLower(strings.TrimSpace(user.Username)) != "user" && user.Password != "pass" {
>>>>>>> 8e4ec4c41d89c9406d3c186dddc3e1129455dab6
		w.WriteHeader(http.StatusForbidden)
		fmt.Println("Error logging in")
		fmt.Fprint(w, "Invalid credentials")
		return
	}
	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
<<<<<<< HEAD
=======

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error extracting the key")
		mw.Fatal(err)
	} // TODO: Missing error definition
>>>>>>> 8e4ec4c41d89c9406d3c186dddc3e1129455dab6
	tokenString, err := token.SignedString(mw.SignKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		mw.Fatal(err)
	}
	response.JSON(mw.Token{tokenString}, w)
}
