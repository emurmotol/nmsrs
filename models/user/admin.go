package user

import (
	"github.com/zneyrl/nmsrs/env"
)

func SetDefaultUser() {
	_, err := FindByEmail(env.AdminEmail)

	if err != nil {
		if err := CreateAdminUser(env.AdminName, env.AdminEmail, env.AdminPassword); err != nil {
			panic(err)
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
	usr.PhotoIsSet = false
	_, err := usr.Insert()

	if err != nil {
		return err
	}
	return nil
}

func IsAdminUser(id string) bool {
	usr, _ := Find(id)

	if usr.Email == env.AdminEmail {
		return true
	}
	return false
}
