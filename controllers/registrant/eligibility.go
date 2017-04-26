package registrant

import (
	"net/http"

	"github.com/zneyrl/nmsrs/helpers/tpl"
)

func Eligibility(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Eligibility",
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "dashboard", "registrant.create.eligibility", data, funcMap)
}

func UpdateEligibility(w http.ResponseWriter, r *http.Request) {
}
