package search

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
)

func Results(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Results",
	}
	tmpl.Render(w, r, "search", "search.results", data, nil)
}
