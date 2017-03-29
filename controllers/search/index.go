package search

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/shared/tmpl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, "search", "search.index", nil)
}
