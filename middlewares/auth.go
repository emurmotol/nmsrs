package middlewares

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/user"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/urfave/negroni"
)

var (
	privKeyPath = homeDir() + "/.ssh/id_rsa"
	pubKeyPath  = homeDir() + "/.ssh/id_rsa.test.pub" // TODO: File manually added original .pub file is incompatible
	signKey     *rsa.PrivateKey
	verifyKey   *rsa.PublicKey
)

func homeDir() string {
	usr, err := user.Current()
	Fatal(err)
	return usr.HomeDir
}

func Fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitKeys() {
	signBytes, err := ioutil.ReadFile(privKeyPath)
	Fatal(err)

	signKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	Fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	Fatal(err)

	verifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	Fatal(err)
}

func validateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		return verifyKey, nil
	})

	if err == nil {
		if token.Valid {
			next(w, r)
		} else {
			// TODO: Redirect to logout
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
		// TODO: Redirect to login
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Unauthorized access to this resource")
	}
}

func Secure(handler http.HandlerFunc) *negroni.Negroni {
	return negroni.New(
		negroni.HandlerFunc(validateToken),
		negroni.Wrap(handler),
	)
}

func GetToken() string {
	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString, err := token.SignedString(signKey)
	Fatal(err)
	return tokenString
}
