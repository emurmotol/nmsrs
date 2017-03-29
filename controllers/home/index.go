package home

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/shared/tmpl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, "main", "home.index", nil)
}

func Welcome(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, "main", "home.welcome", nil)
}
