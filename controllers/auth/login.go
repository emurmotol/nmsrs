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
	var user models.UserCredentials
	err := decoder.Decode(&user, r.PostForm)

	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprint(w, "Error in request")
		return
	}

	if strings.ToLower(strings.TrimSpace(user.Username)) != "user" && user.Password != "pass" {
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
	response.Json(mw.Token{tokenString}, w)
}
