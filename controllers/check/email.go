package check

import (
	"encoding/json"
	"net/http"

	"github.com/emurmotol/nmsrs/helpers/lang"
	"github.com/emurmotol/nmsrs/models/user"
)

func EmailExists(w http.ResponseWriter, r *http.Request) {
	yes := user.IsEmailTaken(r.URL.Query().Get("email"))

	if !yes {
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

func EmailTaken(w http.ResponseWriter, r *http.Request) {
	yes := user.IsEmailTaken(r.URL.Query().Get("email"))

	if yes {
		js, err := json.Marshal(map[string]string{
			"error": lang.En["EmailTaken"],
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

func EmailTakenOrSameAsOld(w http.ResponseWriter, r *http.Request) {
	sameAsOld, err := user.IsEmailSameAsOld(r.URL.Query().Get("id"), r.URL.Query().Get("email"))

	if err != nil {
		panic(err)
	}

	if !sameAsOld {
		yes := user.IsEmailTaken(r.URL.Query().Get("email"))

		if yes {
			js, err := json.Marshal(map[string]string{
				"error": lang.En["EmailTaken"],
			})

			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusNotFound)
			w.Write(js)
			return
		}
	}
	w.WriteHeader(http.StatusOK)
	return
}
