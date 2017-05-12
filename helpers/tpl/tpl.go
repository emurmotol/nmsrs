package tpl

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/emurmotol/nmsrs/env"
	"github.com/emurmotol/nmsrs/helpers/flash"
	"github.com/emurmotol/nmsrs/helpers/lang"
	"github.com/emurmotol/nmsrs/helpers/str"
	"github.com/emurmotol/nmsrs/middlewares"
	"github.com/emurmotol/nmsrs/models/user"
)

var (
	templates map[string]*template.Template
)

func init() {
	// if templates == nil {
	// 	templates = make(map[string]*template.Template)
	// }
	// parseTemplateDir(env.TemplateParentDir)
}

func ParseAllAndRender(w http.ResponseWriter, layout string, name string, data map[string]interface{}) error {
	tpl, ok := templates[layout+":"+name]
	if !ok {
		return fmt.Errorf(lang.En["TemplateNotFound"], name)
	}

	if err := tpl.ExecuteTemplate(w, layout, data); err != nil {
		return err
	}
	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	return nil
} // TODO: Not used

func Render(w http.ResponseWriter, r *http.Request, layout string, name string, data map[string]interface{}, funcMap template.FuncMap) {
	tplFile := env.TemplateParentDir + env.TemplatePathSeparator + strings.Replace(name, ".", env.TemplatePathSeparator, -1) + env.TemplateExt
	layoutFile := env.TemplateParentDir + env.TemplatePathSeparator + env.TemplateLayoutsDir + env.TemplatePathSeparator + layout + env.TemplateExt

	funcMap["DateForHumans"] = str.DateForHumans
	funcMap["IsAdminUser"] = user.IsAdminUser
	t := template.New(fmt.Sprintf("%s:%s", layout, name)).Funcs(funcMap)
	tpl := template.Must(t.ParseFiles(layoutFile, tplFile))

	data["Lang"] = lang.En
	data["Config"] = env.Config()
	data["Request"] = r
	f, err := flash.Get(r, w)

	if err != nil {
		panic(err)
	}
	data["Flash"] = f
	id := middlewares.GetAuthID(w, r)
	var usr *user.User

	if id != "" {
		usr, _ = user.FindByID(id)
	}
	data["AuthUser"] = usr

	if err := tpl.ExecuteTemplate(w, layout, data); err != nil {
		panic(err)
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
			p := strings.TrimPrefix(filepath.Dir(path), root+env.TemplatePathSeparator)
			// TODO: What if there's no layout page template
			if p == env.TemplateLayoutsDir {
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
			p := strings.TrimPrefix(filepath.Dir(page), root+env.TemplatePathSeparator)
			pageDir := strings.Replace(p, env.TemplatePathSeparator, ".", -1)
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
