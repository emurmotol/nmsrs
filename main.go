package main

import (
	"net/http"

	"github.com/urfave/negroni"
	"gopkg.in/unrolled/render.v1"
)

func layout(l string) *render.Render {
	return render.New(render.Options{
		Directory:  "views",
		Layout:     "layouts/" + l,
		Extensions: []string{".gohtml"},
	})
}

func home(w http.ResponseWriter, r *http.Request) {
	layout("main").HTML(w, http.StatusOK, "home/index", map[string]interface{}{
		"title": "Home",
	})
}

func login(w http.ResponseWriter, r *http.Request) {
	layout("auth").HTML(w, http.StatusOK, "auth/login", map[string]interface{}{
		"title": "Login",
	})
}

func register(w http.ResponseWriter, r *http.Request) {
	layout("auth").HTML(w, http.StatusOK, "auth/register", map[string]interface{}{
		"title": "Register",
	})
}

func main() {
	m := http.NewServeMux()

	m.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("static"))))
	m.HandleFunc("/", home)
	m.HandleFunc("/login", login)
	m.HandleFunc("/register", register)

	n := negroni.Classic()
	n.UseHandler(m)
	n.Run(":8080")
}
