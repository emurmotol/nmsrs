package tmpl

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var templates map[string]*template.Template

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	parseTemplateDir("views") // TODO: Add to config for custom folder name
}

func Render(w http.ResponseWriter, layout string, name string, data interface{}) error {
	tmpl, ok := templates[layout+":"+name]
	if !ok {
		return fmt.Errorf("The template %s does not exist", name)
	}

	err := tmpl.ExecuteTemplate(w, layout, data)
	if err != nil {
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
			p := strings.TrimPrefix(filepath.Dir(path), root+string(os.PathSeparator))
			// TODO: Add to config for custom folder name
			// TODO: What if there are no layout page template
			if p == "layouts" {
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
			p := strings.TrimPrefix(filepath.Dir(page), root+string(os.PathSeparator))
			pageDir := strings.Replace(p, string(os.PathSeparator), ".", -1)
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
