package user

import (
	"errors"
	"time"

	"github.com/zneyrl/nmsrs-lookup/db"
	"github.com/zneyrl/nmsrs-lookup/helpers/str"
	"gopkg.in/mgo.v2/bson"
)

type ResetPassword struct {
	Password        string `schema:"password" json:"password" bson:"password,omitempty" validate:"required,min=6"`
	ConfirmPassword string `schema:"confirm_password" json:"confirm_password" bson:",omitempty" validate:"required,eqfield=Password"`
	UpdatedAt       int64  `schema:"updated_at" json:"updated_at" bson:"updatedAt,omitempty"`
}

func UpdatePassword(id string, resetPassword ResetPassword) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("invalid object ID")
	}

	if err := CheckAdmin(id); err != nil {
		return err
	}
	resetPassword.Password = str.Bcrypt(resetPassword.Password)
	resetPassword.ConfirmPassword = ""
	resetPassword.UpdatedAt = time.Now().Unix()

	if err := db.Users.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": resetPassword}); err != nil {
		return err
	}
	return nil
}
