package search

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
)

func Results(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Results",
	}
	tmpl.RenderWithFunc(w, "search", "search.results", data, nil)
}
