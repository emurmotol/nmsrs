package home

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Home",
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "main", "home.index", data, funcMap)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Welcome",
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "main", "home.welcome", data, funcMap)
}
