package search

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/shared/tmpl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Search",
	}
	tmpl.Render(w, "search", "search.index", data)
}
