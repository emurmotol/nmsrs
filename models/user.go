package models

import (
	"errors"
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
	ID              bson.ObjectId `schema:"id" json:"id" bson:"_id"`
	Name            string        `schema:"name" json:"name" bson:"name" validate:"required,min=2"`
	Email           string        `schema:"email" json:"email" bson:"email" validate:"required,email"`
	Password        string        `schema:"password" json:"password" bson:"password" validate:"required,min=6"`
	ConfirmPassword string        `schema:"confirm_password" json:"confirm_password" bson:"confirmPassword" validate:"required,eqfield=Password"` // TODO: Lol
	CreatedAt       int64         `schema:"created_at" json:"created_at" bson:"createdAt"`
	UpdatedAt       int64         `schema:"updated_at" json:"updated_at" bson:"updatedAt"`
}

func (usr *User) All() ([]User, error) {
	usrs := []User{}
	if err := db.Users.Find(bson.M{}).All(&usrs); err != nil {
		return nil, err
	}
	return usrs, nil
}

func (usr *User) Insert() error {
	usr.ID = bson.NewObjectId()
	usr.Password = str.Bcrypt(usr.Password)
	now := time.Now().Unix()
	usr.CreatedAt = now
	usr.UpdatedAt = now

	if err := db.Users.Insert(usr); err != nil {
		return err
	}
	return nil
}

func (usr User) Find(id string) (User, error) {
	if !bson.IsObjectIdHex(id) {
		return usr, errors.New("invalid object id")
	}

	if err := db.Users.FindId(bson.ObjectIdHex(id)).One(&usr); err != nil {
		return usr, err
	}
	return usr, nil
}

func (usr *User) Update(id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("invalid object id")
	}
	usr.ID = bson.ObjectIdHex(id) // TODO: What happened here?
	usr.UpdatedAt = time.Now().Unix()

	if err := db.Users.UpdateId(usr.ID, usr); err != nil {
		return err
	}
	return nil
}

func (usr *User) Delete() error {
	if err := db.Users.RemoveId(usr.ID); err != nil {
		return err
	}
	return nil
}

func (usr User) DeleteMany(ids []string) error {
	for _, id := range ids {
		usr, err := usr.Find(id)

		if err != nil {
			return err
		}

		if err := db.Users.RemoveId(usr.ID); err != nil {
			return err
		}
	}
	return nil
}

func (usr *User) CheckEmailIfTaken() error {
	count, _ := db.Users.Find(bson.M{"email": usr.Email}).Count()

	if count != 0 {
		return errors.New("Email already taken")
	}
	return nil
}

func (usr *User) CheckEmailIfSameAsOld(id string) error {
	u, err := usr.Find(id)

	if err != nil {
		return err
	}

	if u.Email != usr.Email {
		return usr.CheckEmailIfTaken()
	}
	return nil
}
