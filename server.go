package main

import (
	"net/http"

	"github.com/urfave/negroni"
	"github.com/zneyrl/nmsrs-lookup/routes"
)

func main() {
	r := routes.Web()

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(r)
	http.ListenAndServe(":8080", n)
}
