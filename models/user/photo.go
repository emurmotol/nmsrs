package user

import (
	"fmt"
	"mime/multipart"
	"path"
	"path/filepath"
	"strings"

	"github.com/zneyrl/nmsrs-lookup/db"
	"github.com/zneyrl/nmsrs-lookup/env"
	"github.com/zneyrl/nmsrs-lookup/helpers/fi"
	"gopkg.in/mgo.v2/bson"
)

func SetPhoto(file multipart.File, handler *multipart.FileHeader, id string) error {
	filename := fmt.Sprintf("default%s", strings.ToLower(filepath.Ext(handler.Filename)))
	name := filepath.Join(contentDir, id, "photo", filename)

	if err := fi.Save(file, handler, name); err != nil {
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
