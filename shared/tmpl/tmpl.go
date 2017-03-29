package tmpl

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var (
	templates map[string]*template.Template
	// bufpool   *bpool.BufferPool
)

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}
	parseTemplateDir("views") // TODO: Add to config
	// bufpool = bpool.NewBufferPool(64)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Render(w http.ResponseWriter, name string, data interface{}) error {
	tmpl, ok := templates[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist", name)
	}

	// buf := bufpool.Get()
	// defer bufpool.Put(buf)

	err := tmpl.ExecuteTemplate(w, name, data)
	// err := tmpl.ExecuteTemplate(buf, name, data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "text/html")
	// buf.WriteTo(w)
	return nil
}

func paths(root string) ([]string, []string, error) {
	var layouts, pages []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			if removeRootInPath(path, root) == "layouts" { // TODO: Add to config
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
			pageDir := strings.Replace(removeRootInPath(page, root), string(os.PathSeparator), ".", -1)
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

func removeRootInPath(path string, root string) string {
	return strings.TrimPrefix(filepath.Dir(path), root+string(os.PathSeparator))
}
