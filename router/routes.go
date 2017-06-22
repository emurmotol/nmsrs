package router

import (
	"github.com/emurmotol/nmsrs/controller"
	"github.com/pressly/chi"
)

func userRoutes() chi.Router {
	r := chi.NewRouter()
	r.Use(adminOnly)
	r.Get("/", controller.GetUsers)
	r.Post("/", controller.StoreUser)
	r.Get("/create", controller.CreateUser)
	r.Post("/delete", controller.DeleteManyUser)
	r.Get("/email/taken", controller.UserEmailTaken)
	r.Route("/:userId", func(r chi.Router) {
		r.Use(userCtx)
		r.Get("/", controller.ShowUser)
		r.Post("/", controller.UpdateProfile)
		r.Get("/edit", controller.EditUser)
		r.Get("/photo", controller.UserPhoto)
		r.Post("/delete", controller.DeleteUser)
		r.Post("/password/reset", controller.UserPasswordReset)
		r.Get("/email/check", controller.UserEmailCheck)
	})
	return r
}

// func registrantRoutes() chi.Router {
// 	r := chi.NewRouter()
// 	r.Use(adminOnly)
// 	r.Get("/", controller.GetRegistrants)
// 	r.Post("/", controller.StoreRegistrant)
// 	r.Get("/create", controller.CreateRegistrant)
// 	// r.Post("/delete", controller.DeleteManyRegistrant)
// 	r.Get("/email/taken", controller.RegistrantEmailTaken)
// 	// r.Route("/:registrantId", func(r chi.Router) {
// 	// 	r.Use(registrantCtx)
// 	// 	r.Get("/", controller.ShowRegistrant)
// 	// 	r.Post("/", controller.UpdateRegistrant)
// 	// 	r.Get("/edit", controller.EditRegistrant)
// 	// 	r.Get("/photo", controller.RegistrantPhoto)
// 	// 	r.Post("/delete", controller.DeleteRegistrant)
// 	// 	r.Get("/email/check", controller.RegistrantEmailCheck)
// 	// })
// 	return r
// }

func apiRoutes() chi.Router {
	r := chi.NewRouter()
	r.Use(adminOnly)
	// r.Get("/search", controller.SearchIndex)
	r.Get("/certificates", controller.CertificateIndex)
	r.Get("/citymuns/provinces", controller.CityMunProvinceIndex)
	r.Get("/citymuns/:cityMunId/barangays", controller.CityMunBarangayIndex)
	r.Get("/countries", controller.CountryIndex)
	r.Get("/courses", controller.CourseIndex)
	r.Get("/edulevels", controller.EduLevelIndex)
	r.Get("/eligibilities", controller.EligibilityIndex)
	r.Get("/industries", controller.IndustryIndex)
	r.Get("/languages", controller.LanguageIndex)
	r.Get("/licenses", controller.LicenseIndex)
	r.Get("/otherskills", controller.OtherSkillIndex)
	r.Get("/positions", controller.PositionIndex)
	r.Get("/religions", controller.ReligionIndex)
	r.Get("/schools", controller.SchoolIndex)
	r.Get("/skills", controller.SkillIndex)
	r.Get("/unempstats", controller.UnEmpStatIndex)
	return r
}

func profileRoutes() chi.Router {
	r := chi.NewRouter()
	r.Route("/:userId", func(r chi.Router) {
		r.Use(userCtx)
		r.Get("/", controller.ShowUserProfile)
		r.Get("/edit", controller.EditUserProfile)
	})
	return r
}
