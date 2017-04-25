package middlewares

import (
	"net/http"

	"github.com/urfave/negroni"
)

var common = negroni.New(negroni.HandlerFunc(validateToken))

func Admin(handler http.HandlerFunc) *negroni.Negroni {
	return common.With(negroni.Wrap(handler))
}

func Auth(handler http.HandlerFunc) *negroni.Negroni {
	return common.With(negroni.Wrap(handler))
}

func Web(handler http.HandlerFunc) *negroni.Negroni {
	return negroni.New(negroni.Wrap(handler))
}
