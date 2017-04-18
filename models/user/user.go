package user

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/zneyrl/nmsrs-lookup/db"
	"github.com/zneyrl/nmsrs-lookup/helpers/str"
)

type User struct {
	ID              bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Name            string        `schema:"name" json:"name" bson:"name,omitempty" validate:"required,min=2"`
	Email           string        `schema:"email" json:"email" bson:"email,omitempty" validate:"required,email"`
	Password        string        `schema:"password" json:"password" bson:"password,omitempty" validate:"required,min=6"`
	ConfirmPassword string        `schema:"confirm_password" json:"confirm_password" bson:"confirmPassword,omitempty" validate:"required,eqfield=Password"` // TODO: Lol
	CreatedAt       int64         `schema:"created_at" json:"created_at" bson:"createdAt,omitempty"`
	UpdatedAt       int64         `schema:"updated_at" json:"updated_at" bson:"updatedAt,omitempty"`
}

func All() ([]User, error) {
	users := []User{}
	if err := db.Users.Find(bson.M{}).All(&users); err != nil {
		return nil, err
	}
	return users, nil
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

func Find(id string) (User, error) {
	var usr User

	if !bson.IsObjectIdHex(id) {
		return usr, errors.New("Invalid object ID")
	}

	if err := db.Users.FindId(bson.ObjectIdHex(id)).One(&usr); err != nil {
		return usr, err
	}
	return usr, nil
}

func FindByEmail(email string) (User, error) {
	var usr User

	if err := db.Users.Find(bson.M{"email": email}).One(&usr); err != nil {
		return usr, err
	}
	return usr, nil
}

func (usr *User) Delete() error {
	if err := db.Users.RemoveId(usr.ID); err != nil {
		return err
	}
	return nil
}

func DeleteMany(ids []string) error {
	for _, id := range ids {
		usr, err := Find(id)

		if err != nil {
			return err
		}

		if err := db.Users.RemoveId(usr.ID); err != nil {
			return err
		}
	}
	return nil
}

func CheckEmailIfTaken(email string) error {
	count, _ := db.Users.Find(bson.M{"email": email}).Count()

	if count != 0 {
		return errors.New("Email has already been taken")
	}
	return nil
}

func CheckEmailIfSameAsOld(id string, email string) (bool, error) {
	u, err := Find(id)

	if err != nil {
		return false, err
	}

	if u.Email != email {
		return false, nil
	}
	return true, nil
}
