package search

import (
	"net/http"

	"github.com/emurmotol/nmsrs/helpers/tpl"
)

func Results(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Results",
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "search", "search.results", data, funcMap)
}
