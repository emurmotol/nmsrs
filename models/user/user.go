package user

import (
	"errors"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/zneyrl/nmsrs-lookup/db"
	"github.com/zneyrl/nmsrs-lookup/helpers/str"
)

type Src map[string]interface{}

type AuthCredentials struct {
	Email    string `schema:"email" validate:"required,email"`
	Password string `schema:"password" validate:"required"`
}

type User struct {
	ID              bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Name            string        `schema:"name" json:"name" bson:"name,omitempty" validate:"required,min=2"`
	Email           string        `schema:"email" json:"email" bson:"email,omitempty" validate:"required,email"`
	Password        string        `schema:"password" json:"password" bson:"password,omitempty" validate:"required,min=6"`
	ConfirmPassword string        `schema:"confirm_password" json:"confirm_password" bson:"confirmPassword,omitempty" validate:"required,eqfield=Password"` // TODO: Lol
	CreatedAt       int64         `schema:"created_at" json:"created_at" bson:"createdAt,omitempty"`
	UpdatedAt       int64         `schema:"updated_at" json:"updated_at" bson:"updatedAt,omitempty"`
}

type Profile struct {
	Name      string `schema:"name" json:"name" bson:"name,omitempty" validate:"required,min=2"`
	Email     string `schema:"email" json:"email" bson:"email,omitempty" validate:"required,email"`
	UpdatedAt int64  `schema:"updated_at" json:"updated_at" bson:"updatedAt,omitempty"`
}

type ResetPassword struct {
	Password        string `schema:"password" json:"password" bson:"password,omitempty" validate:"required,min=6"`
	ConfirmPassword string `schema:"confirm_password" json:"confirm_password" bson:"confirmPassword,omitempty" validate:"required,eqfield=Password"` // TODO: Lol
	UpdatedAt       int64  `schema:"updated_at" json:"updated_at" bson:"updatedAt,omitempty"`
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
		return usr, errors.New("invalid object id")
	}

	if err := db.Users.FindId(bson.ObjectIdHex(id)).One(&usr); err != nil {
		return usr, err
	}
	return usr, nil
}

func UpdateProfile(id string, profile Profile) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("invalid object id")
	}
	profile.UpdatedAt = time.Now().Unix()

	if err := db.Users.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": profile}); err != nil {
		return err
	}
	return nil
}

func UpdatePassword(id string, resetPassword ResetPassword) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("invalid object id")
	}
	resetPassword.Password = str.Bcrypt(resetPassword.Password)
	resetPassword.UpdatedAt = time.Now().Unix()

	if err := db.Users.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": resetPassword}); err != nil {
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
		return errors.New("Email already taken")
	}
	return nil
}

func CheckEmailIfSameAsOld(id string, email string) error {
	u, err := Find(id)

	if err != nil {
		return err
	}

	if u.Email != email {
		return CheckEmailIfTaken(email)
	}
	return nil
}
