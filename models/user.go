package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/zneyrl/nmsrs-lookup/db"
	"github.com/zneyrl/nmsrs-lookup/helpers/str"
)

type AuthCredentials struct {
	Email    string `schema:"email" validate:"required,email"`
	Password string `schema:"password" validate:"required"`
}

type User struct {
	Name            string `schema:"name" json:"name" bson:"name" validate:"required,min=2"`
	Email           string `schema:"email" json:"email" bson:"email" validate:"required,email"`
	Password        string `schema:"password" json:"password" bson:"password" validate:"required,min=6"`
	ConfirmPassword string `schema:"confirm_password" json:"confirm_password" bson:"confirmPassword" validate:"required,eqfield=Password"` // TODO: Lol
	CreatedAt       int64  `schema:"created_at" json:"created_at" bson:"createdAt"`
	UpdatedAt       int64  `schema:"updated_at" json:"updated_at" bson:"updatedAt"`
}

func (usr *User) All() ([]User, error) {
	usrs := []User{}
	if err := db.Users.Find(bson.M{}).All(&usrs); err != nil {
		return nil, err
	}
	return usrs, nil
}

func (usr *User) Insert() error {
	usr.Password = str.Bcrypt(usr.Password)
	now := time.Now().Unix()
	usr.CreatedAt = now
	usr.UpdatedAt = now

	if err := db.Users.Insert(usr); err != nil {
		return err
	}
	return nil
}

func (usr *User) Find() {

}

func (usr *User) Update() {

}

func (usr *User) Delete() {

}
