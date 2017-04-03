package auth

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/shared/tmpl"
)

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Login",
	}
	tmpl.Render(w, "auth", "auth.login", data)
}
