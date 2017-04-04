package models

type User struct {
	ID        string `schema:"id"`
	Name      string `schema:"name"`
	Username  string `schema:"username"`
	Password  string `schema:"password"`
	CreatedAt string `schema:"created_at"`
	UpdatedAt string `schema:"updated_at"`
}

type AuthCredentials struct {
	Username string `schema:"username" validate:"required"`
	Password string `schema:"password" validate:"required"`
}
