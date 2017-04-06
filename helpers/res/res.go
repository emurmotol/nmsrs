package res

import (
	"encoding/json"
	"net/http"
)

func JSON(response interface{}, w http.ResponseWriter) {
	json, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

type Make struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Errors interface{} `json:"errors"`
}
