package search

import (
	"net/http"

	"github.com/zneyrl/nmsrs/helpers/tmpl"
)

func Results(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Results",
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "search", "search.results", data, funcMap)
}
