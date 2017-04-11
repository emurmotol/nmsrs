package flash

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("secret")) // TODO: Add more secure

func Set(r *http.Request, w http.ResponseWriter, m string) error {
	s, err := Store.Get(r, "flash-session")

	if err != nil {
		return err
	}
	s.AddFlash(m, "flash-message")
	s.Save(r, w)
	return nil
}

func Get(r *http.Request, w http.ResponseWriter) (interface{}, error) {
	s, err := Store.Get(r, "flash-session")

	if err != nil {
		return "", err
	}
	flashes := s.Flashes("flash-message")

	if flashes == nil {
		return nil, nil
	}
	s.Save(r, w)
	return flashes[0], nil
}
