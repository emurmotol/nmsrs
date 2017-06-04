package router

import (
	"net/http"

	"github.com/emurmotol/nmsrs.v4/controller"
	"github.com/emurmotol/nmsrs.v4/env"
	"github.com/emurmotol/nmsrs.v4/helper"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

func Handler() chi.Router {
	r := chi.NewRouter()

	// r.Use(middleware.RequestID)
	// r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(tokenAuth.Verifier)

	r.Get("/", controller.Index)

	r.Group(func(r chi.Router) {
		r.Use(redirectIfAuthenticated)
		r.Use(tokenAuthCtx)
		r.Get("/login", controller.GetLogin)
		r.Post("/login", controller.Login)
		r.Get("/users/email/exists", controller.UserEmailExists)
	})
	r.Get("/logout", controller.Logout)

	staticDir, _ := env.Conf.String("dir.static")
	r.FileServer("/assets", http.Dir(staticDir))

	r.NotFound(controller.NotFound)

	r.Group(func(r chi.Router) {
		r.Use(loggedInOnly)
		r.Mount(helper.ApiBasePath(""), apiRoutes())
		r.Mount("/users", userRoutes())
		r.Mount("/", profileRoutes())
		r.Get("/search", controller.GetSearch)
	})
	return r
}
