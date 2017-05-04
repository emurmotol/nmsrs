package search

import (
	"net/http"

	"github.com/emurmotol/nmsrs/helpers/tpl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Search",
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "search", "search.index", data, funcMap)
}
