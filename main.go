package main

import (
	"net/http"

	"github.com/gorilla/mux"
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

func showHomePage(w http.ResponseWriter, r *http.Request) {
	layout("main").HTML(w, http.StatusOK, "home/index", map[string]interface{}{
		"title": "Home",
	})
}

func showWelcomePage(w http.ResponseWriter, r *http.Request) {
	layout("main").HTML(w, http.StatusOK, "home/welcome", map[string]interface{}{
		"title": "Welcome",
	})
}

func showLoginForm(w http.ResponseWriter, r *http.Request) {
	layout("auth").HTML(w, http.StatusOK, "auth/login", map[string]interface{}{
		"title": "Login",
	})
}

func login(w http.ResponseWriter, r *http.Request) {

} //TODO

func showRegisterForm(w http.ResponseWriter, r *http.Request) {
	layout("auth").HTML(w, http.StatusOK, "auth/register", map[string]interface{}{
		"title": "Register",
	})
}

func register(w http.ResponseWriter, r *http.Request) {

} //TODO

func showSearchForm(w http.ResponseWriter, r *http.Request) {
	layout("search").HTML(w, http.StatusOK, "search/index", map[string]interface{}{
		"title": "Search",
	})
}

func search(w http.ResponseWriter, r *http.Request) {

} //TODO

func showResultsPage(w http.ResponseWriter, r *http.Request) {
	layout("search").HTML(w, http.StatusOK, "search/results", map[string]interface{}{
		"title": "Results",
	})
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/", showHomePage).Methods("GET")
	r.HandleFunc("/welcome", showWelcomePage).Methods("GET")

	r.HandleFunc("/login", showLoginForm).Methods("GET")
	r.HandleFunc("/login", login).Methods("POST") //TODO

	r.HandleFunc("/register", showRegisterForm).Methods("GET")
	r.HandleFunc("/register", register).Methods("POST") //TODO

	r.HandleFunc("/search", showSearchForm).Methods("GET")
	r.HandleFunc("/search", search).Methods("POST") //TODO

	r.HandleFunc("/results", showResultsPage).Methods("GET")

	h := http.StripPrefix("/assets/", http.FileServer(http.Dir("./static/")))
	r.PathPrefix("/assets/").Handler(h).Methods("GET")

	n := negroni.Classic()
	n.UseHandler(r)
	n.Run(":8080")
}

// Password Reset Routes Used By Laravel...
// $this->get('password/reset', 'Auth\ForgotPasswordController@showLinkRequestForm');
// $this->post('password/email', 'Auth\ForgotPasswordController@sendResetLinkEmail');
// $this->get('password/reset/{token}', 'Auth\ResetPasswordController@showResetForm');
// $this->post('password/reset', 'Auth\ResetPasswordController@reset');
