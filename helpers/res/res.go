package res

import (
	"encoding/json"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, response interface{}) {
	json, err := json.Marshal(response)
	if err != nil {
		log.Fatal(err)
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
