package user

import (
	"strings"
	"time"

	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/helpers/str"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

func UpdateProfile(id string, usr User) error {
	if !bson.IsObjectIdHex(id) {
		return models.ErrInvalidObjectID
	}

	if IsAdminUser(id) {
		return ErrActionNotPermitted
	}
	usr.Email = strings.ToLower(usr.Email)
	usr.UpdatedAt = time.Now().Unix()

	if err := db.Users.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": usr}); err != nil {
		return err
	}
	return nil
}

func UpdatePassword(id string, usr User) error {
	if !bson.IsObjectIdHex(id) {
		return models.ErrInvalidObjectID
	}

	if IsAdminUser(id) {
		return ErrActionNotPermitted
	}
	usr.Password = str.Bcrypt(usr.Password)
	usr.UpdatedAt = time.Now().Unix()

	if err := db.Users.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": usr}); err != nil {
		return err
	}
	return nil
}
