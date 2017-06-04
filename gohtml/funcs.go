package gohtml

import (
	"html/template"

	"github.com/emurmotol/nmsrs.v4/env"
	"github.com/emurmotol/nmsrs.v4/helper"
	"github.com/emurmotol/nmsrs.v4/lang"
)

func Funcs() template.FuncMap {
	return template.FuncMap{
		"apiBasePath":   helper.ApiBasePath,
		"baseURL":       helper.BaseURL,
		"config":        env.Map,
		"lang":          lang.Get,
		"dateForHumans": helper.DateForHumans,
		"photoPath":     helper.PhotoPath,
	}
}
