package user

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

func FindByEmail(email string) (*User, error) {
	var usr User

	if err := db.Users.Find(bson.M{"email": email}).One(&usr); err != nil {
		return &usr, err
	}
	return &usr, nil
}

func IsEmailSameAsOld(id string, email string) (bool, error) {
	u, err := FindByID(id)

	if err != nil {
		return false, err
	}

	if u.Email != email {
		return false, nil
	}
	return true, nil
}

func IsEmailTaken(email string) bool {
	count, _ := db.Users.Find(bson.M{"email": email}).Count()

	if count != 0 {
		return true
	}
	return false
}
