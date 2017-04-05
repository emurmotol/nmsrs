package models

type User struct {
	ID              string `schema:"id"`
	Name            string `schema:"name" validate:"required"`
	Username        string `schema:"username" validate:"required,min=10"`
	Password        string `schema:"password" validate:"required"`
	ConfirmPassword string `schema:"confirmPassword" validate:"required"`
	CreatedAt       string `schema:"created_at"`
	UpdatedAt       string `schema:"updated_at"`
}

type AuthCredentials struct {
	Username string `schema:"username" validate:"required"`
	Password string `schema:"password" validate:"required"`
}
