package applicant

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Applicants",
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "dashboard", "applicant.index", data, funcMap)
}
