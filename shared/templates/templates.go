package templates

import (
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo"
)

var templates map[string]*template.Template

type Template struct {
	templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist.", name)
	}
	return tmpl.ExecuteTemplate(w, name, data)
}

func Load() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	parseTemplateDir("views")
}

func paths(root string) ([]string, []string, error) {
	var layouts, pages []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if removeRootInPath(path, root) == "layouts" {
				layouts = append(layouts, path)
			} else {
				pages = append(pages, path)
			}
		}
		return nil
	})

	if err != nil {
		return nil, nil, err
	}
	return layouts, pages, nil
}

func parseTemplateDir(root string) error {
	layouts, pages, err := paths(root)
	if err != nil {
		return err
	}

	var files []string

	for _, layout := range layouts {
		for _, page := range pages {
			files := append(files, layout, page)
			layoutWithExt := filepath.Base(layout)
			layoutNoExt := strings.TrimSuffix(layoutWithExt, filepath.Ext(layoutWithExt))
			pageDir := strings.Replace(removeRootInPath(page, root), string(os.PathSeparator), ".", -1)
			pageWithExt := filepath.Base(page)
			pageNoExt := strings.TrimSuffix(pageWithExt, filepath.Ext(pageWithExt))
			name := layoutNoExt + ":" + pageDir + "." + pageNoExt
			templates[name] = template.Must(template.New(pageNoExt).ParseFiles(files...))
		}
	}
	return nil
}

func removeRootInPath(path string, root string) string {
	return strings.TrimPrefix(filepath.Dir(path), root+string(os.PathSeparator))
}
