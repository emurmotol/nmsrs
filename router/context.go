package router

import (
	"context"
	"net/http"
	"strconv"

	"github.com/emurmotol/nmsrs/constant"
	"github.com/emurmotol/nmsrs/controller"
	"github.com/emurmotol/nmsrs/model"
	"github.com/pressly/chi"
)

func userCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.Atoi(chi.URLParam(r, "userID"))

		if err != nil {
			controller.NotFound(w, r)
			return
		}
		user := model.UserByID(int64(id))

		if user == nil {
			controller.NotFound(w, r)
			return
		}
		ctx := context.WithValue(r.Context(), constant.UserCtxKey, user)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// func registrantCtx(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		id, err := strconv.Atoi(chi.URLParam(r, "registrantID"))

// 		if err != nil {
// 			controller.NotFound(w, r)
// 			return
// 		}
// 		registrant, err := model.RegistrantByID(int64(id))

// 		if err != nil {
// 			controller.NotFound(w, r)
// 			return
// 		}
// 		ctx := context.WithValue(r.Context(), constant.RegistrantCtxKey, registrant)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

func tokenAuthCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), constant.TokenAuthCtxKey, tokenAuth)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
