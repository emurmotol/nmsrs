package user

import (
	"github.com/zneyrl/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

func FindByEmail(email string) (User, error) {
	var usr User

	if err := db.Users.Find(bson.M{"email": email}).One(&usr); err != nil {
		return usr, err
	}
	return usr, nil
}

func CheckEmailIfTaken(email string) error {
	count, _ := db.Users.Find(bson.M{"email": email}).Count()

	if count != 0 {
		return ErrEmailTaken
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
