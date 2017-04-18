package dashboard

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Dashboard",
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "main", "dashboard.index", data, funcMap)
}
