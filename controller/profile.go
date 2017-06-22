package controller

import (
	"net/http"

	"github.com/emurmotol/nmsrs/constant"
	"github.com/emurmotol/nmsrs/model"
)

func ShowUserProfile(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["title"] = "Your Profile"
	data["user"] = r.Context().Value(constant.UserCtxKey).(*model.User)
	data["authUser"] = authUser(r)
	rd.HTML(w, http.StatusOK, "profile/show", data)
}

func EditUserProfile(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value(constant.UserCtxKey).(*model.User)
	authUser := authUser(r)

	if user.Id.Hex() == authUser.Id.Hex() {
		data := make(map[string]interface{})
		data["title"] = "Edit Your Profile"
		data["user"] = user
		data["authUser"] = authUser
		rd.HTML(w, http.StatusOK, "profile/edit", data)
		return
	}
	Forbidden(w, r)
}
