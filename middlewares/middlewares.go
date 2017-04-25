package middlewares

import (
	"net/http"

	"github.com/urfave/negroni"
)

var Common = negroni.New(negroni.HandlerFunc(ValidateToken))

func Admin(handler http.HandlerFunc) *negroni.Negroni {
	return Common.With(negroni.Wrap(handler))
}

func Auth(handler http.HandlerFunc) *negroni.Negroni {
	return Common.With(negroni.Wrap(handler))
}

func Web(handler http.HandlerFunc) *negroni.Negroni {
	return negroni.New(negroni.Wrap(handler))
}
