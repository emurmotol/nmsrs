package auth

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	mw "github.com/zneyrl/nmsrs-lookup/middlewares"
	"github.com/zneyrl/nmsrs-lookup/shared/response"
	"github.com/zneyrl/nmsrs-lookup/shared/tmpl"
)

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Login",
	}
	tmpl.Render(w, "auth", "auth.login", data)
}

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() // TODO: Must handle error

	// var user models.UserCredentials

	// err := json.NewDecoder(r.Body).Decode(&user)

	// decoder := schema.NewDecoder()
	// err := decoder.Decode(user, r.PostForm)

	// if err != nil {
	// 	w.WriteHeader(http.StatusForbidden)
	// 	fmt.Fprint(w, "Error in request")
	// 	return
	// }

	// if strings.ToLower(user.Username) != "user" {
	// 	if user.Password != "pass" {
	// 		w.WriteHeader(http.StatusForbidden)
	// 		fmt.Println("Error logging in")
	// 		fmt.Fprint(w, "Invalid credentials")
	// 		return
	// 	}
	// }

	if strings.ToLower(r.FormValue("username")) != "user" {
		if r.FormValue("password") != "pass" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Println("Error logging in")
			fmt.Fprint(w, "Invalid credentials")
			return
		}
	}
	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims

	// TODO: Missing error definition, comment instead
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)
	// 	fmt.Fprintln(w, "Error extracting the key")
	// 	mw.Fatal(err)
	// }
	tokenString, err := token.SignedString(mw.SignKey)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Error while signing the token")
		mw.Fatal(err)
	}
	response.Json(mw.Token{tokenString}, w)
}
