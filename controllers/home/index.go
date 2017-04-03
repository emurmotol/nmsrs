package home

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/shared/tmpl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Home",
	}
	tmpl.Render(w, "main", "home.index", data)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Welcome",
	}
	tmpl.Render(w, "main", "home.welcome", data)
}
