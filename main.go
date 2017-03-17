package main

import (
	"net/http"

	"github.com/urfave/negroni"
	"gopkg.in/unrolled/render.v1"
)

func main() {
	r := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".gohtml", ".html"},
	})
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		r.HTML(w, http.StatusOK, "example", "World")
	})

	n := negroni.Classic()
	n.UseHandler(mux)
	n.Run(":8080")
}
