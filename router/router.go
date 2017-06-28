package router

import (
	"net/http"

	"github.com/emurmotol/nmsrs/controller"
	"github.com/emurmotol/nmsrs/env"
	"github.com/pressly/chi"
	"github.com/pressly/chi/middleware"
)

func Handler() chi.Router {
	r := chi.NewRouter()

	// r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
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

	r.Group(func(r chi.Router) {
		r.Use(loggedInOnly)
		r.Mount("/api", apiRoutes())
		r.Mount("/users", userRoutes())
		r.Mount("/registrants", registrantRoutes())
		r.Mount("/", profileRoutes())
		r.Get("/search", controller.GetSearch)
	})
	r.NotFound(controller.NotFound)
	return r
}
