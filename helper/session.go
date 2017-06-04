package helper

import (
	"net/http"

	"github.com/emurmotol/nmsrs/env"
	"github.com/gorilla/sessions"
)

func GetSession(r *http.Request, name string) *sessions.Session {
	secret, _ := env.Conf.String("pkg.sessions.secret")
	store := sessions.NewCookieStore([]byte(secret))
	session, err := store.Get(r, name)

	if err != nil {
		panic(err)
	}
	return session
}

func GetFlash(w http.ResponseWriter, r *http.Request, name string) interface{} {
	flash, _ := env.Conf.String("pkg.sessions.flash")
	session := GetSession(r, flash)
	flashes := session.Flashes(name)

	if flashes == nil {
		return nil
	}
	session.Save(r, w)
	return flashes[0]
}

func SetFlash(w http.ResponseWriter, r *http.Request, name string, value interface{}) {
	flash, _ := env.Conf.String("pkg.sessions.flash")
	session := GetSession(r, flash)
	session.AddFlash(value, name)
	session.Save(r, w)
}
