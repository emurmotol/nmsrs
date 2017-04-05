package main

import (
	"net/http"

	"github.com/urfave/negroni"
	"github.com/zneyrl/nmsrs-lookup/middlewares"
	"github.com/zneyrl/nmsrs-lookup/routes"
)

func main() {
	middlewares.InitKeys()
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewRecovery())
	r := routes.Web()
	n.UseHandler(r)
	http.ListenAndServe(":8080", n)
}
