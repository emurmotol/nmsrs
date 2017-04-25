package auth

import (
	"net/http"

	"github.com/zneyrl/nmsrs/env"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:   env.JWTTokenName,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})
	http.Redirect(w, r, env.URL("/login"), http.StatusFound)
}
