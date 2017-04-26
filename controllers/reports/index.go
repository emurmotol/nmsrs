package reports

import (
	"net/http"

	"github.com/zneyrl/nmsrs/helpers/tpl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Reports",
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "menu", "reports.index", data, funcMap)
}
