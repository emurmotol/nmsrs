package models

import (
	"log"
	"time"

	uuid "github.com/satori/go.uuid"
	"github.com/zneyrl/nmsrs-lookup/db"
	"github.com/zneyrl/nmsrs-lookup/helpers/str"
)

var (
	ID  = "_design/users"
	Rev string
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

func init() {
	// TODO: Check if view exsist then do
	// WTF
	rev, err := db.Con.Save(map[string]interface{}{"testing": "Testing"}, ID, "")

	if err != nil {
		log.Fatal(err)
	}
	Rev = rev
}

// func (u *User) All() []User {
// 	// db.Con.
// }

func (u *User) Insert() error {
	u.Password = str.Bcrypt(u.Password)
	now := time.Now().Unix()
	u.CreatedAt = now
	u.UpdatedAt = now
	// WTF
	_, err := db.Con.Save(u, uuid.NewV4().String(), Rev) // TODO: Revision not used

	if err != nil {
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
