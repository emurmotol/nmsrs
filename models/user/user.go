package user

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"gopkg.in/mgo.v2/bson"

	"strings"

	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/env"
	"github.com/emurmotol/nmsrs/helpers/lang"
	"github.com/emurmotol/nmsrs/helpers/str"
	"github.com/emurmotol/nmsrs/models"
)

var (
	ErrActionNotPermitted = errors.New(lang.En["action_not_permitted"])
	ErrEmailTaken         = errors.New(lang.En["email_taken"])
	contentDir            = "content/users"
)

type User struct {
	ID              bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Name            string        `schema:"name" json:"name" bson:"name,omitempty" validate:"required,min=2"`
	Email           string        `schema:"email" json:"email" bson:"email,omitempty" validate:"required,email"`
	Password        string        `schema:"password" json:"password" bson:"password,omitempty" validate:"required,min=6"`
	ConfirmPassword string        `schema:"confirm_password" json:"confirm_password" bson:",omitempty" validate:"required,eqfield=Password"`
	IsAdmin         bool          `schema:"is_admin" json:"is_admin" bson:"isAdmin"`
	PhotoIsSet      bool          `schema:"photo_is_set" json:"photo_is_set" bson:"photoIsSet"`
	CreatedAt       int64         `schema:"created_at" json:"created_at" bson:"createdAt,omitempty"`
	UpdatedAt       int64         `schema:"updated_at" json:"updated_at" bson:"updatedAt,omitempty"`
}

func All() ([]User, error) {
	users := []User{}

	if err := db.Users.Find(bson.M{"email": bson.M{"$ne": env.AdminEmail}}).Sort("+name").All(&users); err != nil {
		return nil, err
	}
	return users, nil
}

func (usr *User) Insert() (string, error) {
	usr.ID = bson.NewObjectId()
	usr.Email = strings.ToLower(usr.Email)
	usr.ConfirmPassword = ""
	usr.Password = str.Bcrypt(usr.Password)
	now := time.Now().Unix()
	usr.CreatedAt = now
	usr.UpdatedAt = now

	if err := db.Users.Insert(usr); err != nil {
		return "", err
	}
	return usr.ID.Hex(), MakeReadMeFile(usr)
}

func FindByID(id string) (*User, error) {
	var usr User

	if !bson.IsObjectIdHex(id) {
		return &usr, models.ErrInvalidObjectID
	}

	if err := db.Users.FindId(bson.ObjectIdHex(id)).One(&usr); err != nil {
		return &usr, err
	}
	return &usr, nil
}

func (usr *User) Delete() error {
	id := usr.ID.Hex()

	if IsAdminUser(id) {
		return ErrActionNotPermitted
	}

	if err := db.Users.RemoveId(usr.ID); err != nil {
		return err
	}
	dir := filepath.Join(contentDir, id)
	_, err := os.Stat(dir)

	if err == nil {
		if err := os.RemoveAll(dir); err != nil {
			return err
		}
	}
	return nil
}

func DeleteMany(ids []string) error {
	for _, id := range ids {
		usr, err := FindByID(id)

		if err != nil {
			return err
		}

		if err := usr.Delete(); err != nil {
			return err
		}
	}
	return nil
} // TODO: Slow on big data, use db.Users.RemoveAll instead

func MakeReadMeFile(usr *User) error {
	file := filepath.Join(contentDir, usr.ID.Hex(), "README.md")

	dir := filepath.Dir(file)
	_, err := os.Stat(dir)

	if os.IsNotExist(err) {
		os.MkdirAll(dir, 0777)
	}
	content := fmt.Sprintf("ID: %s", usr.ID.Hex())
	content += fmt.Sprintf("\nJoined: %s", str.DateForHumans(usr.CreatedAt))

	if err := ioutil.WriteFile(file, []byte(content), 0644); err != nil {
		return err
	}
	return nil
}
