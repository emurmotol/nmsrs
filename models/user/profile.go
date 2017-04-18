package user

import (
	"errors"
	"time"

	"github.com/zneyrl/nmsrs-lookup/db"
	"gopkg.in/mgo.v2/bson"
)

type Profile struct {
	Name      string `schema:"name" json:"name" bson:"name,omitempty" validate:"required,min=2"`
	Email     string `schema:"email" json:"email" bson:"email,omitempty" validate:"required,email"`
	UpdatedAt int64  `schema:"updated_at" json:"updated_at" bson:"updatedAt,omitempty"`
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
