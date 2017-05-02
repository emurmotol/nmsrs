package user

import (
	"time"

	"github.com/zneyrl/nmsrs/db"
	"github.com/zneyrl/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Profile struct {
	Name    string `schema:"name" json:"name" bson:"name,omitempty" validate:"required,min=2"`
	Email   string `schema:"email" json:"email" bson:"email,omitempty" validate:"required,email"`
	IsAdmin bool   `schema:"is_admin" json:"is_admin" bson:"isAdmin"`
	// TODO: PhotoIsSet is handled when user.SetPhoto func is called
	UpdatedAt int64 `schema:"updated_at" json:"updated_at" bson:"updatedAt,omitempty"`
}

func UpdateProfile(id string, profile Profile) error {
	if !bson.IsObjectIdHex(id) {
		return models.ErrInvalidObjectID
	}

	if IsAdminUser(id) {
		return ErrActionNotPermitted
	}
	profile.UpdatedAt = time.Now().Unix()

	if err := db.Users.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": profile}); err != nil {
		return err
	}
	return nil
}
