package registrant

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
)

func CertificateOfCompetence(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Certificate Of Competence",
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "dashboard", "registrant.create.certificate_of_competence", data, funcMap)
}

func UpdateCertificateOfCompetence(w http.ResponseWriter, r *http.Request) {
}
