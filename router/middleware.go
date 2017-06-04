package router

import (
	"context"
	"net/http"

	"github.com/emurmotol/nmsrs.v4/constant"
	"github.com/emurmotol/nmsrs.v4/controller"
	"github.com/emurmotol/nmsrs.v4/env"
	"github.com/emurmotol/nmsrs.v4/model"
	"github.com/goware/jwtauth"
)

var (
	tokenAuth *jwtauth.JwtAuth
)

func init() {
	secret, _ := env.Conf.String("pkg.jwtauth.secret")
	tokenAuth = jwtauth.New("HS256", []byte(secret), nil)
}

func adminOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user := r.Context().Value(constant.AuthUserCtxKey).(*model.User)
		if !user.IsAdmin {
			controller.Forbidden(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func loggedInOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user, err := model.GetAuthorizedUser(r)

		if err != nil {
			http.Redirect(w, r, "/login?redirect="+r.URL.Path, http.StatusFound)
			return
		}

		if user == nil {
			controller.NotFound(w, r)
			return
		}
		ctx := context.WithValue(r.Context(), constant.AuthUserCtxKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func redirectIfAuthenticated(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if user, _ := model.GetAuthorizedUser(r); user != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		next.ServeHTTP(w, r)
	})
}
