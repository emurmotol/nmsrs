package controllers

import (
	"net/http"

	"github.com/zneyrl/nmsrs/helpers/tpl"
)

func PageNotFound(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Page not found",
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "error", "error.404", data, funcMap)
}
