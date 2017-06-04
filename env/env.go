package env

import (
	"io/ioutil"

	"github.com/olebedev/config"
)

var Conf *config.Config

func init() {
	yml, err := ioutil.ReadFile("env/conf.yml")

	if err != nil {
		panic(err)
	}
	cfg, err := config.ParseYaml(string(yml))

	if err != nil {
		panic(err)
	}
	Conf = cfg
}

func Map() map[string]interface{} {
	mp, err := Conf.Map("")

	if err != nil {
		panic(err)
	}
	return mp
}
