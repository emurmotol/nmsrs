package main

import (
	"net/http"

	"github.com/urfave/negroni"
	"github.com/zneyrl/nmsrs-lookup/routes"
	"github.com/zneyrl/nmsrs-lookup/middlewares"
)

func main() {
	middlewares.InitKeys()

	r := routes.Web()

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewRecovery())
	n.UseHandler(r)
	http.ListenAndServe(":8080", n)
}
