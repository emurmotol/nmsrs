package search

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Search",
	}
	tmpl.RenderWithFunc(w, "search", "search.index", data, nil)
}
