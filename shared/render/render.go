package render

import (
	"html/template"
	"io"

	"github.com/labstack/echo"
)

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Dir(dir string) {
	t := &Template{
		templates: template.Must(template.ParseGlob("views/" + dir + "/*.gohtml")),
	}

	// e.Renderer = t
}
