package registrant

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
)

func WorkExperience(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Work Experience",
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "dashboard", "registrant.create.work_experience", data, funcMap)
}

func UpdateWorkExperience(w http.ResponseWriter, r *http.Request) {
}
