package controller

import (
	"net/http"

	"github.com/emurmotol/nmsrs/model"
)

func Index(w http.ResponseWriter, r *http.Request) {
	if user, _ := model.GetAuthorizedUser(r); user != nil {
		data := make(map[string]interface{})
		data["auth"] = user
		dashboard(w, r, data)
		return
	}
	welcome(w, r)
}

func welcome(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["title"] = "Welcome"
	data["urlPath"] = r.URL.Path
	rd.HTML(w, http.StatusOK, "home/welcome", data)
}

func dashboard(w http.ResponseWriter, r *http.Request, data map[string]interface{}) {
	data["title"] = "Dashboard"
	data["urlPath"] = r.URL.Path
	rd.HTML(w, http.StatusOK, "home/dashboard", data)
}
