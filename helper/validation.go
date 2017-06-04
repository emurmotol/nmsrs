package helper

import (
	"github.com/emurmotol/nmsrs/env"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	translation "gopkg.in/go-playground/validator.v9/translations/en"
)

var (
	validate   *validator.Validate
	uni        *ut.UniversalTranslator
	translator ut.Translator
)

func initValidatorVars() {
	locale := en.New()
	uni = ut.New(locale, locale)
	lang, _ := env.Conf.String("app.lang")
	translator, _ = uni.GetTranslator(lang)
	validate = validator.New()
	translation.RegisterDefaultTranslations(validate, translator)
}

func ValidateForm(s interface{}) map[string]string {
	initValidatorVars()

	err := validate.Struct(s)
	errs := make(map[string]string)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			panic(err)
		}

		for _, e := range err.(validator.ValidationErrors) {
			errs[e.Field()] = e.Translate(translator) // todo: displayed field should not be in camel case
		}
	}
	return errs
}
