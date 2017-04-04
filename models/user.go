package models

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

<<<<<<< HEAD
type AuthCredentials struct {
=======
type UserCredentials struct {
>>>>>>> 8e4ec4c41d89c9406d3c186dddc3e1129455dab6
	Username string `schema:"username"`
	Password string `schema:"password"`
}
