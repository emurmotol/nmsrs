package client

import (
	"log"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/zneyrl/nmsrs-lookup/env"
	"github.com/zneyrl/nmsrs-lookup/middlewares"
)

func GetToken(id string) string {
	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["id"] = id
	claims[env.JWTTokenName] = "level1"                                   // TODO: WTF level1 means
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(168)).Unix() // TODO: Expires in 1 week
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString, err := token.SignedString(middlewares.SignKey)

	if err != nil {
		log.Fatal(err)
	}
	return tokenString
}

func GetAuthID(w http.ResponseWriter, r *http.Request) string {
	tokenCookie, err := r.Cookie(env.JWTTokenName)

	if err == nil {
		token, err := jwt.Parse(tokenCookie.Value, func(token *jwt.Token) (interface{}, error) {
			return middlewares.VerifyKey, nil
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
