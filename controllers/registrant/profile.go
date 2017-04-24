package registrant

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Profile",
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "dashboard", "registrant.create.profile", data, funcMap)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
}
