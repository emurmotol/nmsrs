package dashboard

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Dashboard",
	}
	tmpl.RenderWithFunc(w, "dashboard", "dashboard.index", data, nil)
}

func Overview(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Overview",
	}
	tmpl.RenderWithFunc(w, "main", "dashboard.overview", data, nil)
}
