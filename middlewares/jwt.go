package middlewares

import (
	"crypto/rsa"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/zneyrl/nmsrs/env"
	"github.com/zneyrl/nmsrs/helpers/lang"
)

var (
	signKey   *rsa.PrivateKey
	verifyKey *rsa.PublicKey
)

func init() {
	signBytes, err := ioutil.ReadFile(env.KeyPrivate)
	if err != nil {
		panic(err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		panic(err)
	}

	verifyBytes, err := ioutil.ReadFile(env.KeyPublic)
	if err != nil {
		panic(err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		panic(err)
	}
}

func validateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	tokenCookie, err := r.Cookie(env.JWTTokenName)

	switch {
	case err == http.ErrNoCookie:
		http.Redirect(w, r, env.URL("/login"), http.StatusFound)
		return
	case err != nil:
		panic(err)
	}

	token, err := jwt.Parse(tokenCookie.Value, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
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
			panic(err)
		}
	default:
		panic(err)
	}
}

func GetToken(id string) string {
	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["id"] = id
	claims[env.JWTTokenName] = "level1"                                   // TODO: WTF level1 means
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(168)).Unix() // TODO: Expires in 1 week
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString, err := token.SignedString(signKey)

	if err != nil {
		panic(err)
	}
	return tokenString
}

func GetAuthID(r *http.Request) string {
	tokenCookie, err := r.Cookie(env.JWTTokenName)

	if err == nil {
		token, err := jwt.Parse(tokenCookie.Value, func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})

		if err != nil {
			// TODO: Logout to fix this error
			panic(err)
		}
		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok && !token.Valid {
			panic(lang.En["token_invalid"])
		}

		if claims["id"] != nil {
			return claims["id"].(string)
		}
		// http.Redirect(w, r, env.URL("/logout"), http.StatusFound)
	}
	// ErrNoCookie, No auth user
	return ""
}
