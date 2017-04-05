package auth

import (
	"net/http"
	"strings"
	"time"

	validator "gopkg.in/go-playground/validator.v9"

	jwt "github.com/dgrijalva/jwt-go"
	ut "github.com/go-playground/universal-translator"
	"github.com/gorilla/schema"
	mw "github.com/zneyrl/nmsrs-lookup/middlewares"
	"github.com/zneyrl/nmsrs-lookup/models"
	"github.com/zneyrl/nmsrs-lookup/shared/res"
	"github.com/zneyrl/nmsrs-lookup/shared/str"
	"github.com/zneyrl/nmsrs-lookup/shared/tmpl"
	"github.com/zneyrl/nmsrs-lookup/shared/trans"
)

var (
	decoder  = schema.NewDecoder()
	validate *validator.Validate
	uni      *ut.UniversalTranslator
)

func init() {
	en, _ := uni.GetTranslator("en")
}

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Login",
	}
	tmpl.Render(w, "auth", "auth.login", data)
}

func Login(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		res.JSON(res.Make{http.StatusInternalServerError, "", "Error parsing form"}, w)
		return
	}
	var user models.AuthCredentials
	err = decoder.Decode(&user, r.PostForm)

	if err != nil {
		res.JSON(res.Make{http.StatusInternalServerError, "", "Error in request"}, w)
		return
	}
	validate = validator.New()
	err = validate.Struct(user)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			res.JSON(res.Make{http.StatusInternalServerError, "", err}, w)
			return
		}
		errs := make(map[string]string)

		for _, err := range err.(validator.ValidationErrors) {
			errs[str.LowerCaseFirstChar(err.Field())] = trans.GetEq(str.CamelCaseToSentenceCase(err.Field()), err.Tag())
		}
		res.JSON(res.Make{http.StatusForbidden, "", errs}, w)
		return
	} // TODO: Create a package for this

	if strings.ToLower(user.Username) != "user" || user.Password != "pass" {
		res.JSON(res.Make{http.StatusForbidden, "", "Invalid credentials"}, w)
		return
	}
	token := jwt.New(jwt.SigningMethodRS256)
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString, err := token.SignedString(mw.SignKey)

	if err != nil {
		res.JSON(res.Make{http.StatusInternalServerError, "", "Error while signing the token"}, w)
		mw.Fatal(err)
	}
	// TODO: Redirect to dashboard
	res.JSON(res.Make{http.StatusOK, map[string]string{
		"token":   tokenString,
		"message": "Success login!",
	}, ""}, w)
	return
}
