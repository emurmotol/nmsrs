package models

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type AuthCredentials struct {
	Username string `schema:"username" validate:"required,email"`
	Password string `schema:"password" validate:"required"`
}
