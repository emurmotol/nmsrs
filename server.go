package main

import (
	"net/http"

	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
	"github.com/zneyrl/nmsrs-lookup/shared/tmpl"
)

func main() {
	p := pat.New()
	p.Get("/", func(w http.ResponseWriter, r *http.Request) {
		tmpl.Render(w, "search", "search.index", nil)
	})

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(p)
	http.ListenAndServe(":8080", n)
}
