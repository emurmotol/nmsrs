package models

type AuthCredentials struct {
	Email    string `schema:"email" validate:"required,email"`
	Password string `schema:"password" validate:"required"`
}

type User struct {
	Id              string `schema:"id"`
	Name            string `schema:"name" validate:"required,min=2"`
	Email           string `schema:"email" validate:"required,email"`
	Password        string `schema:"password" validate:"required,min=6"`
	ConfirmPassword string `schema:"confirmPassword" validate:"required,eqfield=Password"` // TODO: Lol
	CreatedAt       string `schema:"createdAt"`
	UpdatedAt       string `schema:"updatedAt"`
}

func (u *User) Insert() {

}

func (u *User) Find() {

}

func (u *User) Update() {

}

func (u *User) Delete() {

}
