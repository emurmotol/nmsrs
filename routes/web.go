package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zneyrl/nmsrs-lookup/controllers/auth"
	"github.com/zneyrl/nmsrs-lookup/controllers/dashboard"
	"github.com/zneyrl/nmsrs-lookup/controllers/home"
	"github.com/zneyrl/nmsrs-lookup/controllers/search"
	mw "github.com/zneyrl/nmsrs-lookup/middlewares"
)

func Web() *mux.Router {
	r := mux.NewRouter()
	r.Path("/").Methods("GET").Handler(mw.Secure(home.Index))
	r.Path("/welcome").Methods("GET").HandlerFunc(home.Welcome)

	login := r.Path("/login").Subrouter()
	login.Methods("GET").HandlerFunc(auth.ShowLoginForm)
	login.Methods("POST").HandlerFunc(auth.Login)

	register := r.Path("/register").Subrouter()
	register.Methods("GET").HandlerFunc(auth.ShowRegisterForm)
	register.Methods("POST").HandlerFunc(auth.Register)

	r.Path("/dashboard").Methods("GET").Handler(mw.Secure(dashboard.Index))

	r.Path("/search").Methods("GET").HandlerFunc(search.Index)
	r.Path("/results").Methods("GET").HandlerFunc(search.Results)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	return r
}
