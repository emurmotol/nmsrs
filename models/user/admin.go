package user

import (
	"errors"
	"log"

	"github.com/zneyrl/nmsrs-lookup/env"
)

func SetDefaultUser() {
	_, err := FindByEmail(env.AdminEmail)

	if err != nil {
		if err := CreateAdminUser(env.AdminName, env.AdminEmail, env.AdminPassword); err != nil {
			log.Fatal(err)
		}
	}
}

func CreateAdminUser(name string, email string, password string) error {
	var usr User
	usr.Name = name
	usr.Email = email
	usr.Password = password
	usr.ConfirmPassword = ""
	usr.IsAdmin = true

	if err := usr.Insert(); err != nil {
		return err
	}
	return nil
}

func CheckAdmin(id string) error {
	usr, _ := Find(id)

	if usr.Email == env.AdminEmail {
		return errors.New("Action not permitted")
	}
	return nil
}
