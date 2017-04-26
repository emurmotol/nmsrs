package user

import (
	"mime/multipart"
	"path"
	"path/filepath"

	"github.com/zneyrl/nmsrs/db"
	"github.com/zneyrl/nmsrs/env"
	"github.com/zneyrl/nmsrs/helpers/img"
	"gopkg.in/mgo.v2/bson"
)

// var photoPath =

func SetPhoto(file multipart.File, id string) error {
	if IsAdminUser(id) {
		return ErrActionNotPermitted
	}
	name := filepath.Join(contentDir, id, "photo", filepath.Base(env.DefaultUserPhoto))

	if err := img.SaveAsJPEG(file, name); err != nil {
		return err
	}

	if err := db.Users.UpdateId(bson.ObjectIdHex(id), bson.M{"$set": bson.M{"photoIsSet": true}}); err != nil {
		return err
	}
	return nil
}

func GetPhoto(id string) string {
	return path.Join(contentDir, id, "photo", filepath.Base(env.DefaultUserPhoto))
}
