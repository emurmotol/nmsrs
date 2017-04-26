package registrant

import (
	"net/http"

	"github.com/zneyrl/nmsrs/helpers/tpl"
)

func FormalEducation(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Formal Education",
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "dashboard", "registrant.create.formal_education", data, funcMap)
}

func UpdateFormalEducation(w http.ResponseWriter, r *http.Request) {
}