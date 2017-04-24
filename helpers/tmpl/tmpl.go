package tmpl

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/zneyrl/nmsrs-lookup/env"
	"github.com/zneyrl/nmsrs-lookup/helpers/flash"
	"github.com/zneyrl/nmsrs-lookup/helpers/str"
	"github.com/zneyrl/nmsrs-lookup/middlewares"
	"github.com/zneyrl/nmsrs-lookup/models/user"
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

func ParseAllAndRender(w http.ResponseWriter, layout string, name string, data map[string]interface{}) error {
	tmpl, ok := templates[layout+":"+name]
	if !ok {
		return fmt.Errorf("The template %s does not exist", name)
	}

	if err := tmpl.ExecuteTemplate(w, layout, data); err != nil {
		return err
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	return nil
} // TODO: Not used

func Render(w http.ResponseWriter, r *http.Request, layout string, name string, data map[string]interface{}, funcMap template.FuncMap) {
	tmplFile := layoutsParentDir + pathSeparator + strings.Replace(name, ".", pathSeparator, -1) + tmplExt
	layoutFile := layoutsParentDir + pathSeparator + layoutsDir + pathSeparator + layout + tmplExt

	funcMap["DateForHumans"] = str.DateForHumans
	t := template.New(fmt.Sprintf("%s:%s", layout, name)).Funcs(funcMap)
	tmpl := template.Must(t.ParseFiles(layoutFile, tmplFile))

	data["Config"] = env.Config()
	data["Request"] = r
	f, err := flash.Get(r, w)

	if err != nil {
		log.Fatal(err)
	}
	data["Flash"] = f
	id := middlewares.GetAuthUserID(w, r)
	var usr user.User

	if id != "" {
		usr, _ = user.Find(id)
	}
	data["AuthUser"] = usr

	if err := tmpl.ExecuteTemplate(w, layout, data); err != nil {
		log.Fatal(err)
	}
	w.Header().Set("Content-Type", "text/html")
}

func paths(root string) ([]string, []string, error) {
	var layouts, pages []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			p := strings.TrimPrefix(filepath.Dir(path), root+pathSeparator)
			// TODO: What if there's no layout page template
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
} // TODO: Not used

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
} // TODO: Not used
