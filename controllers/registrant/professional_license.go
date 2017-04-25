package registrant

import (
	"net/http"

	"github.com/zneyrl/nmsrs/helpers/tmpl"
)

func ProfessionalLicense(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Professional License",
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "dashboard", "registrant.create.professional_license", data, funcMap)
}

func UpdateProfessionalLicense(w http.ResponseWriter, r *http.Request) {
}
