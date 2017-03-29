package home

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/shared/templates"
)

func Index(w http.ResponseWriter, r *http.Request) {
	templates.Render(w, "main:home.index", nil)
}
