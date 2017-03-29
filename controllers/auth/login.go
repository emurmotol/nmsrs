package auth

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/shared/tmpl"
)

func ShowLoginForm(w http.ResponseWriter, r *http.Request) {
	tmpl.Render(w, "auth", "auth.login", nil)
}
