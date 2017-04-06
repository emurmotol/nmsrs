package reports

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Reports",
		"R":     r,
	}
	tmpl.RenderWithFunc(w, "dashboard", "reports.index", data, nil)
}
