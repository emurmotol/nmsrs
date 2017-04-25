package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"github.com/zneyrl/nmsrs-lookup/controllers/auth"
	"github.com/zneyrl/nmsrs-lookup/controllers/check"
	"github.com/zneyrl/nmsrs-lookup/controllers/home"
	"github.com/zneyrl/nmsrs-lookup/controllers/registrant"
	"github.com/zneyrl/nmsrs-lookup/controllers/reports"
	"github.com/zneyrl/nmsrs-lookup/controllers/search"
	"github.com/zneyrl/nmsrs-lookup/controllers/user"
)

func Register() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// Auth routes
	authRoutes := mux.NewRouter().StrictSlash(true)
	authRoutes.Path("/logout").Methods("GET").HandlerFunc(auth.Logout)
	authRoutes.Path("/check/file/image/{id}").Methods("POST").HandlerFunc(check.Image)
	authRoutes.Path("/reports").Methods("GET").HandlerFunc(reports.Index)
	authRoutes.Path("/search").Methods("GET").HandlerFunc(search.Index)
	authRoutes.Path("/results").Methods("GET").HandlerFunc(search.Results)

	// Admin routes
	adminRoutes := mux.NewRouter().StrictSlash(true)
	users := adminRoutes.PathPrefix("/users").Subrouter()
	users.Path("/").Methods("GET").HandlerFunc(user.Index)
	users.Path("/create").Methods("GET").HandlerFunc(user.Create)
	users.Path("/").Methods("POST").HandlerFunc(user.Store)
	users.Path("/ids").Methods("POST").HandlerFunc(user.DestroyMany)
	users.Path("/{id}").Methods("GET").HandlerFunc(user.Show)
	users.Path("/{id}/edit").Methods("GET").HandlerFunc(user.Edit)
	users.Path("/{id}/photo").Methods("GET").HandlerFunc(user.Photo)
	users.Path("/{id}").Methods("PUT").HandlerFunc(user.UpdateProfile)
	users.Path("/{id}").Methods("DELETE").HandlerFunc(user.Destroy)
	users.Path("/{id}/reset-password").Methods("POST").HandlerFunc(user.ResetPassword)
	registrants := adminRoutes.PathPrefix("/registrants").Subrouter()
	registrants.Path("/").Methods("GET").HandlerFunc(registrant.Index)
	registrants.Path("/").Methods("POST").HandlerFunc(registrant.Store)
	registrants.Path("/ids").Methods("POST").HandlerFunc(registrant.DestroyMany)
	registrants.Path("/{id}").Methods("GET").HandlerFunc(registrant.Show)
	registrants.Path("/{id}/photo").Methods("GET").HandlerFunc(registrant.Photo)
	registrants.Path("/{id}").Methods("DELETE").HandlerFunc(registrant.Destroy)
	registrants.Path("/{id}/profile").Methods("GET").HandlerFunc(registrant.Profile)
	registrants.Path("/{id}/profile").Methods("PUT").HandlerFunc(registrant.UpdateProfile)
	registrants.Path("/{id}/formal-education").Methods("GET").HandlerFunc(registrant.FormalEducation)
	registrants.Path("/{id}/formal-education").Methods("PUT").HandlerFunc(registrant.UpdateFormalEducation)
	registrants.Path("/{id}/professional-license").Methods("GET").HandlerFunc(registrant.ProfessionalLicense)
	registrants.Path("/{id}/professional-license").Methods("PUT").HandlerFunc(registrant.UpdateProfessionalLicense)
	registrants.Path("/{id}/eligibility").Methods("GET").HandlerFunc(registrant.Eligibility)
	registrants.Path("/{id}/eligibility").Methods("PUT").HandlerFunc(registrant.UpdateEligibility)
	registrants.Path("/{id}/technical-training-and-relevant-experience").Methods("GET").HandlerFunc(registrant.TechnicalTrainingAndRelevantExperience)
	registrants.Path("/{id}/technical-training-and-relevant-experience").Methods("PUT").HandlerFunc(registrant.UpdateTechnicalTrainingAndRelevantExperience)
	registrants.Path("/{id}/certificate-of-competence").Methods("GET").HandlerFunc(registrant.CertificateOfCompetence)
	registrants.Path("/{id}/certificate-of-competence").Methods("PUT").HandlerFunc(registrant.UpdateCertificateOfCompetence)
	registrants.Path("/{id}/work-experience").Methods("GET").HandlerFunc(registrant.WorkExperience)
	registrants.Path("/{id}/work-experience").Methods("PUT").HandlerFunc(registrant.UpdateWorkExperience)
	registrants.Path("/{id}/other-skills-aquired-without-formal-training").Methods("GET").HandlerFunc(registrant.OtherSkillsAquiredWithoutFormalTraining)
	registrants.Path("/{id}/other-skills-aquired-without-formal-training").Methods("PUT").HandlerFunc(registrant.UpdateOtherSkillsAquiredWithoutFormalTraining)
	registrants.Path("/{id}/certification-authorization").Methods("GET").HandlerFunc(registrant.CertificationAuthorization)
	registrants.Path("/{id}/certification-authorization").Methods("PUT").HandlerFunc(registrant.UpdateCertificationAuthorization)

	// Web routes
	webRoutes := mux.NewRouter().StrictSlash(true)
	webRoutes.Path("/").Methods("GET").HandlerFunc(home.Index)
	webRoutes.Path("/welcome").Methods("GET").HandlerFunc(home.Welcome)
	login := webRoutes.PathPrefix("/login").Subrouter()
	login.Methods("GET").HandlerFunc(auth.ShowLoginForm)
	login.Methods("POST").HandlerFunc(auth.Login)

	// common := negroni.New(negroni.HandlerFunc(middlewares.ValidateToken))
	// router.PathPrefix("/auth").Handler(common.With(negroni.Wrap(authRoutes)))
	// router.PathPrefix("/admin").Handler(common.With(negroni.Wrap(adminRoutes)))
	router.PathPrefix("/").Handler(negroni.New(negroni.Wrap(webRoutes)))
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	return router
}
