package auth

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/helpers/tmpl"
)

func ShowRegisterForm(w http.ResponseWriter, r *http.Request) {
	data := map[string]interface{}{
		"Title": "Register",
	}
	funcMap := map[string]interface{}{}
	tmpl.Render(w, r, "auth", "auth.register", data, funcMap)
}

func Register(w http.ResponseWriter, r *http.Request) {
	// TODO: Logic as user store
}
