package home

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Home",
	}
	tmpl.Render(w, r, "main", "home.index", data, nil)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Welcome",
	}
	tmpl.Render(w, r, "main", "home.welcome", data, nil)
}
