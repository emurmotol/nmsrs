package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zneyrl/nmsrs-lookup/controllers/applicant"
	"github.com/zneyrl/nmsrs-lookup/controllers/auth"
	"github.com/zneyrl/nmsrs-lookup/controllers/check"
	"github.com/zneyrl/nmsrs-lookup/controllers/dashboard"
	"github.com/zneyrl/nmsrs-lookup/controllers/home"
	"github.com/zneyrl/nmsrs-lookup/controllers/reports"
	"github.com/zneyrl/nmsrs-lookup/controllers/search"
	"github.com/zneyrl/nmsrs-lookup/controllers/user"
	"github.com/zneyrl/nmsrs-lookup/middlewares"
)

func Web() *mux.Router {
	route := mux.NewRouter().StrictSlash(true)
	route.Path("/").Methods("GET").HandlerFunc(home.Index)
	route.Path("/welcome").Methods("GET").HandlerFunc(home.Welcome)

	login := route.Path("/login").Subrouter()
	login.Methods("GET").HandlerFunc(auth.ShowLoginForm)
	login.Methods("POST").HandlerFunc(auth.Login)

	route.Path("/logout").Methods("GET").Handler(middlewares.Secure(auth.Logout))

	route.Path("/dashboard").Methods("GET").Handler(middlewares.Secure(dashboard.Index))

	route.Path("/users").Methods("GET").Handler(middlewares.Secure(user.Index))
	route.Path("/users/create").Methods("GET").Handler(middlewares.Secure(user.Create))
	route.Path("/users").Methods("POST").Handler(middlewares.Secure(user.Store))
	route.Path("/users/ids").Methods("POST").Handler(middlewares.Secure(user.DestroyMany))
	route.Path("/users/{id}").Methods("GET").Handler(middlewares.Secure(user.Show))
	route.Path("/users/{id}/edit").Methods("GET").Handler(middlewares.Secure(user.Edit))
	route.Path("/users/{id}/photo").Methods("GET").Handler(middlewares.Secure(user.Photo))
	route.Path("/users/{id}").Methods("PUT").Handler(middlewares.Secure(user.UpdateProfile))
	route.Path("/users/{id}").Methods("DELETE").Handler(middlewares.Secure(user.Destroy))
	route.Path("/users/{id}/reset-password").Methods("POST").Handler(middlewares.Secure(user.ResetPassword))

	route.Path("/applicants").Methods("GET").Handler(middlewares.Secure(applicant.Index))
	route.Path("/reports").Methods("GET").Handler(middlewares.Secure(reports.Index))

	route.Path("/check/file/image/{id}").Methods("POST").Handler(middlewares.Secure(check.Image))

	route.Path("/search").Methods("GET").Handler(middlewares.Secure(search.Index))
	route.Path("/results").Methods("GET").Handler(middlewares.Secure(search.Results))
	route.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	return route
}
