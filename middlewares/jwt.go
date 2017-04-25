package middlewares

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/zneyrl/nmsrs/env"
)

var (
	SignKey   *rsa.PrivateKey
	VerifyKey *rsa.PublicKey
)

func init() {
	signBytes, err := ioutil.ReadFile(env.KeyPrivate)
	if err != nil {
		log.Fatal(err)
	}

	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal(err)
	}

	verifyBytes, err := ioutil.ReadFile(env.KeyPublic)
	if err != nil {
		log.Fatal(err)
	}

	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal(err)
	}
}

func ValidateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	tokenCookie, err := r.Cookie(env.JWTTokenName)

	switch {
	case err == http.ErrNoCookie:
		http.Redirect(w, r, env.URL("/login"), http.StatusFound)
		return
	case err != nil:
		log.Fatal(err)
	}

	token, err := jwt.Parse(tokenCookie.Value, func(token *jwt.Token) (interface{}, error) {
		return VerifyKey, nil
	})

	switch err.(type) {
	case nil:
		if !token.Valid {
			http.Redirect(w, r, env.URL("/logout"), http.StatusFound)
			return
		}
		next(w, r)
		return
	case *jwt.ValidationError:
		validationError := err.(*jwt.ValidationError)

		switch validationError.Errors {
		case jwt.ValidationErrorExpired:
			http.Redirect(w, r, env.URL("/login"), http.StatusFound)
			return
		default:
			log.Fatal(err)
		}
	default:
		log.Fatal(err)
	}
}
