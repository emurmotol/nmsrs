package user

import (
	"errors"
	"time"

	"github.com/zneyrl/nmsrs-lookup/db"
	"gopkg.in/mgo.v2/bson"
)

type Profile struct {
	Name       string `schema:"name" json:"name" bson:"name,omitempty" validate:"required,min=2"`
	Email      string `schema:"email" json:"email" bson:"email,omitempty" validate:"required,email"`
	IsAdmin    bool   `schema:"is_admin" json:"is_admin" bson:"isAdmin"`
	PhotoIsSet bool   `schema:"photo_is_set" json:"photo_is_set" bson:"photoIsSet"`
	UpdatedAt  int64  `schema:"updated_at" json:"updated_at" bson:"updatedAt,omitempty"`
}

func UpdateProfile(id string, profile Profile) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("Invalid object ID")
	}

	if err := CheckAdmin(id); err != nil {
		return err
	}
	profile.UpdatedAt = time.Now().Unix()

	if err := db.Users.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": profile}); err != nil {
		return err
	}
	return nil
}
