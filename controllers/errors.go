package controllers

import (
	"net/http"

	"github.com/zneyrl/nmsrs/helpers/tmpl"
)

func PageNotFound(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Page Not Found",
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "error", "error.404", data, funcMap)
}
