package registrant

import (
	"net/http"

	"github.com/zneyrl/nmsrs/helpers/tmpl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Registrants",
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "menu", "registrant.index", data, funcMap)
}
