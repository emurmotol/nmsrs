package user

import (
	"time"

	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/helpers/str"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type ResetPassword struct {
	Password        string `schema:"password" json:"password" bson:"password,omitempty" validate:"required,min=6"`
	ConfirmPassword string `schema:"confirm_password" json:"confirm_password" bson:"-" validate:"required,eqfield=Password"`
	UpdatedAt       int64  `schema:"updated_at" json:"updated_at" bson:"updatedAt,omitempty"`
}

func UpdatePassword(id string, resetPassword ResetPassword) error {
	if !bson.IsObjectIdHex(id) {
		return models.ErrInvalidObjectID
	}

	if IsAdminUser(id) {
		return ErrActionNotPermitted
	}
	resetPassword.Password = str.Bcrypt(resetPassword.Password)
	resetPassword.UpdatedAt = time.Now().Unix()

	if err := db.Users.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": resetPassword}); err != nil {
		return err
	}
	return nil
}
