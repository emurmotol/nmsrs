package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/zneyrl/nmsrs/controllers"
	"github.com/zneyrl/nmsrs/controllers/auth"
	"github.com/zneyrl/nmsrs/controllers/check"
	"github.com/zneyrl/nmsrs/controllers/home"
	"github.com/zneyrl/nmsrs/controllers/registrant"
	"github.com/zneyrl/nmsrs/controllers/reports"
	"github.com/zneyrl/nmsrs/controllers/search"
	"github.com/zneyrl/nmsrs/controllers/user"
	"github.com/zneyrl/nmsrs/middlewares"
)

func Register() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Admin routes
	users := router.PathPrefix("/users").Subrouter()
	users.Path("/create").Methods("GET").Handler(middlewares.Admin(user.Create))
	users.Path("/ids").Methods("POST").Handler(middlewares.Admin(user.DestroyMany))
	users.Path("/{id}/reset-password").Methods("POST").Handler(middlewares.Admin(user.ResetPassword))
	users.Path("/{id}/edit").Methods("GET").Handler(middlewares.Admin(user.Edit))
	users.Path("/{id}/photo").Methods("GET").Handler(middlewares.Admin(user.Photo))
	users.Path("/{id}").Methods("GET").Handler(middlewares.Admin(user.Show))
	users.Path("/{id}").Methods("DELETE").Handler(middlewares.Admin(user.Destroy))
	users.Path("/{id}").Methods("PUT").Handler(middlewares.Admin(user.UpdateProfile))
	users.Path("").Methods("GET").Handler(middlewares.Admin(user.Index))
	users.Path("").Methods("POST").Handler(middlewares.Admin(user.Store))
	registrants := router.PathPrefix("/registrants").Subrouter()
	registrants.Path("/create").Methods("GET").Handler(middlewares.Admin(registrant.Create))
	registrants.Path("/ids").Methods("POST").Handler(middlewares.Admin(registrant.DestroyMany))
	registrants.Path("/{id}/photo").Methods("GET").Handler(middlewares.Admin(registrant.Photo))
	registrants.Path("/{id}").Methods("GET").Handler(middlewares.Admin(registrant.Show))
	registrants.Path("/{id}").Methods("DELETE").Handler(middlewares.Admin(registrant.Destroy))
	registrants.Path("").Methods("GET").Handler(middlewares.Admin(registrant.Index))
	registrants.Path("").Methods("POST").Handler(middlewares.Admin(registrant.Store))

	// Auth routes
	router.Path("/check/file/image/{field}").Methods("POST").Handler(middlewares.Auth(check.Image))
	router.Path("/reports").Methods("GET").Handler(middlewares.Auth(reports.Index))
	router.Path("/search").Methods("GET").Handler(middlewares.Auth(search.Index))
	router.Path("/results").Methods("GET").Handler(middlewares.Auth(search.Results))

	// Web routes
	login := router.PathPrefix("/login").Subrouter()
	login.Path("").Methods("GET").Handler(middlewares.Web(auth.ShowLoginForm))
	login.Path("").Methods("POST").Handler(middlewares.Web(auth.Login))
	router.Path("/logout").Methods("GET").Handler(middlewares.Web(auth.Logout))
	router.Path("/welcome").Methods("GET").Handler(middlewares.Web(home.Welcome))
	router.Path("/").Methods("GET").Handler(middlewares.Web(home.Index))

	router.NotFoundHandler = http.HandlerFunc(controllers.PageNotFound) // TODO: Only works when root/subrouter has path /

	return router
}
