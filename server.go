package main

import (
	"fmt"
	"net/http"

	"github.com/urfave/negroni"
	"github.com/zneyrl/nmsrs/env"
	"github.com/zneyrl/nmsrs/models/certificate"
	"github.com/zneyrl/nmsrs/models/civilstatus"
	"github.com/zneyrl/nmsrs/models/user"
	"github.com/zneyrl/nmsrs/routes"
)

func main() {
	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.Use(negroni.NewRecovery())
	n.Use(negroni.NewStatic(http.Dir("static")))
	n.UseHandler(routes.Register())

	user.SetDefaultUser()

	user.Seeder()
	certificate.Seeder()
	civilstatus.Seeder()

	host := fmt.Sprintf("%s:%d", env.SvrHost, env.SvrPort)
	fmt.Printf("Server running at %s\n", host)

	if err := http.ListenAndServe(host, n); err != nil {
		panic(err)
	}
}
