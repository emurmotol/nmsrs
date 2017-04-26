package registrant

import (
	"net/http"

	"github.com/zneyrl/nmsrs/helpers/tpl"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Profile",
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "dashboard", "registrant.create.profile", data, funcMap)
}

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
}
