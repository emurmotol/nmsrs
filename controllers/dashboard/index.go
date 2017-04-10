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
	tmpl.Render(w, r, "dashboard", "dashboard.index", data, funcMap)
}

func Overview(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Overview",
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "main", "dashboard.overview", data, funcMap)
}
