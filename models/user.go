package models

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

<<<<<<< HEAD
type AuthCredentials struct {
	Username string `schema:"username" validate:"required,email"`
	Password string `schema:"password" validate:"required"`
=======
type UserCredentials struct {
	Username string `schema:"username"`
	Password string `schema:"password"`
>>>>>>> 8e4ec4c41d89c9406d3c186dddc3e1129455dab6
}
