package registrant

import (
	"net/http"

	"github.com/zneyrl/nmsrs/helpers/tpl"
	"github.com/zneyrl/nmsrs/models/registrant"
)

func Create(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title":                   "Create Registrant",
		"EmploymentStatusOptions": registrant.EmploymentStatusOptions,
		"SexOptions":              registrant.SexOptions,
		"CivilStatusOptions":      registrant.CivilStatusOptions,
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "wizard", "registrant.create", data, funcMap)
}

func Store(w http.ResponseWriter, r *http.Request) {
}
