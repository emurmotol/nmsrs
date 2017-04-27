package registrant

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Step(w http.ResponseWriter, r *http.Request) {
	step, err := strconv.Atoi(mux.Vars(r)["step"])

	if err != nil {
		panic(err)
	}

	switch step {
	case 1:
		return
	case 2:
		return
	case 3:
		return
	case 41:
		return
	case 5:
		return
	case 6:
		return
	case 7:
		return
	case 8:
		return
	case 9:
		return
	}
}
