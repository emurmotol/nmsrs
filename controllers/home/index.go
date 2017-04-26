package home

import (
	"net/http"

	"github.com/zneyrl/nmsrs/helpers/tpl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Home",
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "main", "home.index", data, funcMap)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Welcome",
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "main", "home.welcome", data, funcMap)
}
