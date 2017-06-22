package controller

import (
	"html/template"

	"github.com/emurmotol/nmsrs/env"
	"github.com/emurmotol/nmsrs/gohtml"
	"github.com/gorilla/schema"
	"github.com/unrolled/render"
)

var (
	rd *render.Render
	// paginator globals
	limit    int
	interval int
	// schema
	decoder *schema.Decoder
)

func init() {
	decoder = schema.NewDecoder()

	dir, _ := env.Conf.String("pkg.render.dir")
	layout, _ := env.Conf.String("pkg.render.layout")
	ext, _ := env.Conf.String("pkg.render.ext")
	indentJSON, _ := env.Conf.Bool("pkg.render.indentJSON")
	isDev, _ := env.Conf.Bool("pkg.render.isDev")

	rd = render.New(render.Options{
		Directory:     dir,
		Layout:        layout,
		Funcs:         []template.FuncMap{gohtml.Funcs()},
		Extensions:    []string{ext},
		IndentJSON:    indentJSON,
		IsDevelopment: isDev,
	})
	limit, _ = env.Conf.Int("pkg.helper.pagination.limit")
	interval, _ = env.Conf.Int("pkg.helper.pagination.interval")
}
