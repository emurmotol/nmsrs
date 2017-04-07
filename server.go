package main

import (
	"net/http"

	"fmt"

	"github.com/urfave/negroni"
	"github.com/zneyrl/nmsrs-lookup/env"
	mw "github.com/zneyrl/nmsrs-lookup/middlewares"
	"github.com/zneyrl/nmsrs-lookup/routes"
)

func main() {
	mw.InitKeys()
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewRecovery())
	r := routes.Web()
	n.UseHandler(r)
	http.ListenAndServe(fmt.Sprintf("%s:%d", env.SvrHost, env.SvrPort), n)
}
