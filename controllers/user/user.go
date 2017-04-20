package user

import (
	"net/http"

	"github.com/gorilla/schema"
)

var (
	decoder = schema.NewDecoder()
)

func Photo(w http.ResponseWriter, r *http.Request) {

}
