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
	r.Route("/:id", func(r chi.Router) {
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

func apiRoutes() chi.Router {
	r := chi.NewRouter()
	r.Use(adminOnly)
	r.Get("/search", controller.SearchIndex)
	return r
}

func profileRoutes() chi.Router {
	r := chi.NewRouter()
	r.Route("/:id", func(r chi.Router) {
		r.Use(userCtx)
		r.Get("/", controller.ShowUserProfile)
		r.Get("/edit", controller.EditUserProfile)
	})
	return r
}
