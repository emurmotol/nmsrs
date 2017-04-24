package registrant

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
)

func Eligibility(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Eligibility",
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "dashboard", "registrant.create.eligibility", data, funcMap)
}

func UpdateEligibility(w http.ResponseWriter, r *http.Request) {
}
