package registrant

import (
	"net/http"

	"github.com/zneyrl/nmsrs/helpers/tpl"
)

func ProfessionalLicense(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Professional License",
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "dashboard", "registrant.create.professional_license", data, funcMap)
}

func UpdateProfessionalLicense(w http.ResponseWriter, r *http.Request) {
}
