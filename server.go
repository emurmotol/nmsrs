package main

import (
	"net/http"

	"fmt"

	"github.com/urfave/negroni"
	"github.com/zneyrl/nmsrs-lookup/env"
	"github.com/zneyrl/nmsrs-lookup/models/user"
	"github.com/zneyrl/nmsrs-lookup/routes"
)

func main() {
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewRecovery())
	r := routes.Web()
	n.UseHandler(r)
	user.SetDefaultUser()
	http.ListenAndServe(fmt.Sprintf("%s:%d", env.SvrHost, env.SvrPort), n)
}
