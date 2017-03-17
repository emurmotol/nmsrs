package main

import (
	"net/http"

	"github.com/urfave/negroni"
	"gopkg.in/unrolled/render.v1"
)

var (
	rnd = render.New(render.Options{
		Directory:  "views",
		Layout:     "layouts/main",
		Extensions: []string{".gohtml", ".html"},
	})
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "home/index", map[string]interface{}{
		"title": "Home",
	})
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "auth/login", map[string]interface{}{
		"title": "Login",
	})
}

func registerHandler(w http.ResponseWriter, r *http.Request) {
	rnd.HTML(w, http.StatusOK, "auth/register", map[string]interface{}{
		"title": "Register",
	})
}

func main() {
	m := http.NewServeMux()

	m.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("static"))))
	m.HandleFunc("/", homeHandler)
	m.HandleFunc("/login", loginHandler)
	m.HandleFunc("/register", registerHandler)

	n := negroni.Classic()
	n.UseHandler(m)
	n.Run(":8080")
}
