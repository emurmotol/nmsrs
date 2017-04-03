package search

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/shared/tmpl"
)

func Results(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Results",
	}
	tmpl.Render(w, "search", "search.results", data)
}
