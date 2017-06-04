package lang

import (
	"github.com/bykovme/gotrans"
	"github.com/emurmotol/nmsrs/env"
)

func init() {
	if err := gotrans.InitLocales("lang"); err != nil {
		panic(err)
	}
}

func Get(key string) string {
	lang, _ := env.Conf.String("app.lang")
	l := gotrans.DetectLanguage(lang)
	return gotrans.Tr(l, key)
}
