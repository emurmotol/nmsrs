package auth

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/shared/tmpl"
)

func ShowRegisterForm(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Register",
	}
	tmpl.Render(w, "auth", "auth.register", data)
}
