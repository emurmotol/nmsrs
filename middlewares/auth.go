package middlewares

import (
	"crypto/rsa"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/urfave/negroni"
	"github.com/zneyrl/nmsrs-lookup/env"
)

var (
	privateKeyPath = "keys/.ssh/app.rsa"     // openssl genrsa -out app.rsa keysize
	publicKeyPath  = "keys/.ssh/app.rsa.pub" // openssl rsa -in app.rsa -pubout > app.rsa.pub
	signKey        *rsa.PrivateKey
	verifyKey      *rsa.PublicKey
	TokenName      = "AccessToken"
)

func init() {
	signBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Fatal(err)
	}

	verifyBytes, err := ioutil.ReadFile(publicKeyPath)
	if err != nil {
		log.Fatal(err)
	}

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		log.Fatal(err)
	}
}

func validateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	tokenCookie, err := r.Cookie(TokenName)

	switch {
	case err == http.ErrNoCookie:
		http.Redirect(w, r, env.URL("/login"), http.StatusFound)
		return
	case err != nil:
		log.Fatal(err)
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
			log.Fatal(err)
		}
	default:
		log.Fatal(err)
	}
}

func Secure(handler http.HandlerFunc) *negroni.Negroni {
	return negroni.New(
		negroni.HandlerFunc(validateToken),
		negroni.Wrap(handler),
	) // TODO: Understand how this works
}

func GetUserToken(id string) string {
	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["id"] = id
	claims[TokenName] = "level1"                                          // TODO: WTF level1 means
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(168)).Unix() // TODO: Expires in 1 week
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString, err := token.SignedString(signKey)

	if err != nil {
		log.Fatal(err)
	}
	return tokenString
}

func GetAuthUserID(w http.ResponseWriter, r *http.Request) string {
	tokenCookie, err := r.Cookie(TokenName)

	if err == nil {
		token, err := jwt.Parse(tokenCookie.Value, func(token *jwt.Token) (interface{}, error) {
			return verifyKey, nil
		})

		if err != nil {
			log.Fatal(err)
		}
		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok && !token.Valid {
			log.Fatal("invalid JWT token")
		}
		return claims["id"].(string)
	}
	// ErrNoCookie, No auth user
	return ""
}
