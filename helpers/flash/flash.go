package flash

import (
	"net/http"

	"github.com/gorilla/sessions"
	"github.com/emurmotol/nmsrs/env"
)

var (
	Store        = sessions.NewCookieStore([]byte(env.AppKey))
	sessionName  = "flash-session"
	alertVarName = "flash-alert"
)

func Set(r *http.Request, w http.ResponseWriter, txt string) error {
	s, err := Store.Get(r, sessionName)

	if err != nil {
		return err
	}

	s.AddFlash(txt, alertVarName)
	s.Save(r, w)
	return nil
}

func Get(r *http.Request, w http.ResponseWriter) (interface{}, error) {
	s, err := Store.Get(r, sessionName)

	if err != nil {
		return nil, err
	}
	flashes := s.Flashes(alertVarName)

	if len(flashes) == 0 {
		return nil, nil
	}
	s.Save(r, w)
	return flashes[0], nil
}
