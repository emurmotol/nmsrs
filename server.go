package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/pat"
	"github.com/urfave/negroni"
)

func main() {
	p := pat.New()
	p.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// tmpl.Render(w, "main:home.index", nil)

		// t := template.Must(template.New("main").ParseFiles(`views\layouts\main.gohtml`, `views\home\index.gohtml`))
		// err := t.ExecuteTemplate(w, "main", nil)

		t := template.Must(template.New("main:home.index").ParseFiles(`views\layouts\main.gohtml`, `views\home\index.gohtml`))
		err := t.ExecuteTemplate(w, "main:home.index", nil)

		if err != nil {
			log.Fatal(err)
		}
	})

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(p)
	http.ListenAndServe(":8080", n)
}
