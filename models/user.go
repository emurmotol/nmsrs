package models

import (
	"time"

	"github.com/zneyrl/nmsrs-lookup/db"
	"github.com/zneyrl/nmsrs-lookup/helpers/str"
)

type AuthCredentials struct {
	Email    string `schema:"email" validate:"required,email"`
	Password string `schema:"password" validate:"required"`
}

type User struct {
	Name            string `schema:"name" json:"name" validate:"required,min=2"`
	Email           string `schema:"email" json:"email" validate:"required,email"`
	Password        string `schema:"password" json:"password" validate:"required,min=6"`
	ConfirmPassword string `schema:"confirm_password" json:"confirm_password" validate:"required,eqfield=Password"` // TODO: Lol
	CreatedAt       int64  `schema:"created_at" json:"created_at"`
	UpdatedAt       int64  `schema:"updated_at" json:"updated_at"`
}

// func (u *User) All() []User {
// 	// db.Con.
// }

func (u *User) Insert() error {
	u.Password = str.Bcrypt(u.Password)
	now := time.Now().Unix()
	u.CreatedAt = now
	u.UpdatedAt = now

	if err := db.Users.Insert(u); err != nil {
		return err
	}
	return nil
}

func (u *User) Find() {

}

func (u *User) Update() {

}

func (u *User) Delete() {

}
