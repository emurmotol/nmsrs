package models

import (
	"log"

	"github.com/zneyrl/nmsrs-lookup/db"
)

type AuthCredentials struct {
	Email    string `schema:"email" validate:"required,email"`
	Password string `schema:"password" validate:"required"`
}

type User struct {
	ID              string `schema:"_id"`
	Name            string `schema:"name" validate:"required,min=2"`
	Email           string `schema:"email" validate:"required,email"`
	Password        string `schema:"password" validate:"required,min=6"`
	ConfirmPassword string `schema:"confirm_password" validate:"required,eqfield=password"` // TODO: Lol
	CreatedAt       string `schema:"created_at"`
	UpdatedAt       string `schema:"updated_at"`
}

func (u *User) Insert() string {
	rev, err := db.Con.Save(u, "001", "")

	if err != nil {
		log.Fatal(err)
	}
	return rev
}

func (u *User) Find() {

}

func (u *User) Update() {

}

func (u *User) Delete() {

}
