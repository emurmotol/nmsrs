package user

import (
	"net/http"

	"github.com/zneyrl/nmsrs-lookup/shared/tmpl"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Users",
	}
	tmpl.Render(w, "main", "user.index", data)
}

func Create(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Create User",
	}
	tmpl.Render(w, "main", "user.create", data)
}

func Store(w http.ResponseWriter, r *http.Request) {
	// models.User.Insert()
}

func Show(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Show User",
	}
	tmpl.Render(w, "main", "user.show", data)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Title": "Edit User",
	}
	tmpl.Render(w, "main", "user.edit", data)
}

func Update(w http.ResponseWriter, r *http.Request) {
	// models.User.Update()
}

func Destroy(w http.ResponseWriter, r *http.Request) {
	// models.User.Delete()
}
