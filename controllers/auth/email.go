package auth

import (
	"encoding/json"
	"net/http"

	"github.com/emurmotol/nmsrs/helpers/lang"
	"github.com/emurmotol/nmsrs/models/user"
)

func CheckEmailExists(w http.ResponseWriter, r *http.Request) {
	_, err := user.FindByEmail(r.URL.Query().Get("email"))

	if err != nil {
		js, err := json.Marshal(map[string]string{
			"error": lang.En["EmailNotRecognized"],
		})

		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(js)
		return
	}
	w.WriteHeader(http.StatusOK)
	return
}
