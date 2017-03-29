package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zneyrl/nmsrs-lookup/controllers/auth"
	"github.com/zneyrl/nmsrs-lookup/controllers/home"
	"github.com/zneyrl/nmsrs-lookup/controllers/search"
)

func Web() *mux.Router {
	r := mux.NewRouter()
	r.Path("/").Methods("GET").HandlerFunc(home.Index)
	r.Path("/welcome").Methods("GET").HandlerFunc(home.Welcome)
	r.Path("/login").Methods("GET").HandlerFunc(auth.ShowLoginForm)
	r.Path("/register").Methods("GET").HandlerFunc(auth.ShowRegisterForm)
	r.Path("/search").Methods("GET").HandlerFunc(search.Index)
	r.Path("/results").Methods("GET").HandlerFunc(search.Results)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	return r
}
