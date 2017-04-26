package reports

import (
	"net/http"

	"github.com/zneyrl/nmsrs/helpers/tmpl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Reports",
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "menu", "reports.index", data, funcMap)
}
