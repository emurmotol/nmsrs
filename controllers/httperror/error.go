package httperror

import (
	"net/http"

	"github.com/zneyrl/nmsrs/helpers/tmpl"
)

func Err404(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Error 404",
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "error", "error.404", data, funcMap)
}
