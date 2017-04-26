package vald

import (
	en "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/zneyrl/nmsrs/env"
	"github.com/zneyrl/nmsrs/helpers/str"
	validator "gopkg.in/go-playground/validator.v9"
	en_trans "gopkg.in/go-playground/validator.v9/translations/en"
)

var (
	validate *validator.Validate
	uni      *ut.UniversalTranslator
	trans    ut.Translator
)

func init() {
	en := en.New()
	uni = ut.New(en, en)
	trans, _ = uni.GetTranslator(env.Locale)
	validate = validator.New()
	en_trans.RegisterDefaultTranslations(validate, trans)
}

func StructHasError(s interface{}) map[string]string {
	err := validate.Struct(s)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}
		errs := make(map[string]string)

		for _, e := range err.(validator.ValidationErrors) {
			errs[str.CamelCaseToSnakeCase(e.Field())] = str.CamelCaseToSentenceCase(e.Translate(trans)) // TODO: Must parse only the key not the value
		}
		return errs
	}
	return map[string]string{}
}
