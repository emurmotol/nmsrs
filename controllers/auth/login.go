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
)

var decoder = schema.NewDecoder()

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Login",
	}
	tmpl.Render(w, "auth", "auth.login", data)
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // TODO: Must handle error
<<<<<<< HEAD
	var user models.AuthCredentials
=======
	var user models.UserCredentials
>>>>>>> 8e4ec4c41d89c9406d3c186dddc3e1129455dab6
	err := decoder.Decode(&user, r.PostForm)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error in request")
		return
	}

<<<<<<< HEAD
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

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error extracting the key")
		mw.Fatal(err)
	} // TODO: Missing error definition
	tokenString, err := token.SignedString(mw.SignKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		mw.Fatal(err)
	}
	response.JSON(mw.Token{tokenString}, w)
}
