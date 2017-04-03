package middlewares

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/urfave/negroni"
)

var (
	privKeyPath = homeDir() + "/.ssh/id_rsa"
	pubKeyPath  = homeDir() + "/.ssh/id_rsa.manual.pub" // TODO: File manually added
	SignKey     *rsa.PrivateKey
	VerifyKey   *rsa.PublicKey
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

	SignKey, err = jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	Fatal(err)

	verifyBytes, err := ioutil.ReadFile(pubKeyPath)
	Fatal(err)

	VerifyKey, err = jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	Fatal(err)
}

type Token struct {
	Token string `json:"token"`
}

func validateToken(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(token *jwt.Token) (interface{}, error) {
		return VerifyKey, nil
	})

	if err == nil {
		if token.Valid {
			next(w, r)
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprint(w, "Token is not valid")
		}
	} else {
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
