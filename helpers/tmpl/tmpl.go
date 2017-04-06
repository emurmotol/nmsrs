package tmpl

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	templates        map[string]*template.Template
	layoutsParentDir = "views"
	layoutsDir       = "layouts"
	tmplExt          = ".gohtml"
	pathSeparator    = string(os.PathSeparator)
)

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	// parseTemplateDir(layoutsParentDir)
}

func Render(w http.ResponseWriter, layout string, name string, data interface{}) error {
	tmpl, ok := templates[layout+":"+name]
	if !ok {
		return fmt.Errorf("The template %s does not exist", name)
	}

	if err := tmpl.ExecuteTemplate(w, layout, data); err != nil {
		return err
	}
	w.Header().Set("Content-Type", "text/html")
	return nil
}

func RenderWithFunc(w http.ResponseWriter, layout string, name string, data interface{}, funcMap template.FuncMap) error {
	tmplFile := layoutsParentDir + pathSeparator + strings.Replace(name, ".", pathSeparator, -1) + tmplExt
	layoutFile := layoutsParentDir + pathSeparator + layoutsDir + pathSeparator + layout + tmplExt
	t := template.New(fmt.Sprintf("%s:%s", layout, name)).Funcs(funcMap)
	tmpl := template.Must(t.ParseFiles(layoutFile, tmplFile))

	if err := tmpl.ExecuteTemplate(w, layout, data); err != nil {
		return err
	}
	w.Header().Set("Content-Type", "text/html")
	return nil
}

func paths(root string) ([]string, []string, error) {
	var layouts, pages []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			p := strings.TrimPrefix(filepath.Dir(path), root+pathSeparator)
			// TODO: What if there are no layout page template
			if p == layoutsDir {
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

	for _, layout := range layouts {
		for _, page := range pages {
			files := []string{layout, page}
			layoutWithExt := filepath.Base(layout)
			layoutNoExt := strings.TrimSuffix(layoutWithExt, filepath.Ext(layoutWithExt))
			p := strings.TrimPrefix(filepath.Dir(page), root+pathSeparator)
			pageDir := strings.Replace(p, pathSeparator, ".", -1)
			pageWithExt := filepath.Base(page)
			pageNoExt := strings.TrimSuffix(pageWithExt, filepath.Ext(pageWithExt))

			var name string
			if pageDir != root {
				name = layoutNoExt + ":" + pageDir + "." + pageNoExt
			} else {
				name = layoutNoExt + ":" + pageNoExt
			}
			templates[name] = template.Must(template.New(name).ParseFiles(files...))
		}
	}
	return nil
}
