package registrant

import (
	"net/http"

	"github.com/zneyrl/nmsrs/helpers/tpl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Registrants",
	}
	funcMap := map[string]interface{}{}
	tpl.Render(w, r, "menu", "registrant.index", data, funcMap)
}

func Show(w http.ResponseWriter, r *http.Request) {
}
