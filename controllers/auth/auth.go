package auth

import (
	en "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/gorilla/schema"
	validator "gopkg.in/go-playground/validator.v9"
	en_trans "gopkg.in/go-playground/validator.v9/translations/en"
)

var (
	decoder  = schema.NewDecoder()
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	trans    ut.Translator
)

func init() {
	en := en.New()
	uni = ut.New(en, en)
	trans, _ = uni.GetTranslator("en")
	validate = validator.New()
	en_trans.RegisterDefaultTranslations(validate, trans)
}
