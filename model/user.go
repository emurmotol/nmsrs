package model

import (
	"errors"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"mime/multipart"

	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/emurmotol/nmsrs/database"
	"github.com/emurmotol/nmsrs/env"
	"github.com/emurmotol/nmsrs/helper"
	"github.com/emurmotol/nmsrs/lang"
	"github.com/icrowley/fake"
	"golang.org/x/crypto/bcrypt"
)

var (
	errActionNotPermitted = errors.New(lang.Get("action_not_permitted"))
)

type User struct {
	ID        int64      `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	Name      string     `gorm:"not null" json:"name"`
	Email     string     `gorm:"unique;not null" json:"email"`
	Password  string     `gorm:"not null" json:"password"`
	IsAdmin   bool       `gorm:"type:tinyint(1);default:false;not null" json:"is_admin"`
	HasPhoto  bool       `gorm:"type:tinyint(1);default:false;not null" json:"has_photo"`
}

type LoginForm struct {
	Email    string            `schema:"email" validate:"required,email"`
	Password string            `schema:"password" validate:"required"`
	Errors   map[string]string `schema:"-"`
}

func (form *LoginForm) IsValid() bool {
	form.Errors = make(map[string]string)

	if errs := helper.ValidateForm(form); len(errs) != 0 {
		form.Errors = errs
	}

	if taken, _ := UserEmailTaken(form.Email); !taken {
		form.Errors["Email"] = lang.Get("email_not_recognized")
	}
	return len(form.Errors) == 0
}

type CreateUserForm struct {
	Name            string                `schema:"name" validate:"required"`
	Email           string                `schema:"email" validate:"required,email"`
	Password        string                `schema:"password" validate:"required"`
	ConfirmPassword string                `schema:"confirm_password" validate:"eqfield=Password"`
	IsAdmin         bool                  `schema:"is_admin"`
	PhotoFile       multipart.File        `schema:"-"`
	PhotoHeader     *multipart.FileHeader `schema:"-"`
	Errors          map[string]string     `schema:"-"`
}

func (form *CreateUserForm) IsValid() bool {
	form.Errors = make(map[string]string)

	if errs := helper.ValidateForm(form); len(errs) != 0 {
		form.Errors = errs
	}

	if taken, _ := UserEmailTaken(form.Email); taken {
		form.Errors["Email"] = lang.Get("email_taken")
	}

	if form.PhotoFile != nil {
		if err := helper.ValidateImage(form.PhotoHeader); err != nil {
			form.Errors["Photo"] = err.Error()
		}
	}
	return len(form.Errors) == 0
}

type EditProfileForm struct {
	ID          int64                 `schema:"-"`
	Name        string                `schema:"name" validate:"required"`
	Email       string                `schema:"email" validate:"required,email"`
	IsAdmin     bool                  `schema:"is_admin"`
	PhotoFile   multipart.File        `schema:"-"`
	PhotoHeader *multipart.FileHeader `schema:"-"`
	Errors      map[string]string     `schema:"-"`
}

func (form *EditProfileForm) IsValid() bool {
	form.Errors = make(map[string]string)

	if errs := helper.ValidateForm(form); len(errs) != 0 {
		form.Errors = errs
	}

	if same, _ := UserEmailSameAsOld(form.ID, form.Email); !same {
		if taken, _ := UserEmailTaken(form.Email); taken {
			form.Errors["Email"] = lang.Get("email_taken")
		}
	}

	if form.PhotoFile != nil {
		if err := helper.ValidateImage(form.PhotoHeader); err != nil {
			form.Errors["Photo"] = err.Error()
		}
	}
	return len(form.Errors) == 0
}

type PasswordResetForm struct {
	NewPassword     string            `schema:"new_password" validate:"required"`
	ConfirmPassword string            `schema:"confirm_password" validate:"eqfield=NewPassword"`
	Errors          map[string]string `schema:"-"`
}

func (form *PasswordResetForm) IsValid() bool {
	form.Errors = make(map[string]string)

	if errs := helper.ValidateForm(form); len(errs) != 0 {
		form.Errors = errs
	}
	return len(form.Errors) == 0
}

func (user User) Search(q string) []User {
	db := database.Conn()
	defer db.Close()

	users := []User{}
	results := make(chan []User)
	like := database.WrapLike(q)

	go func() {
		db.Not("email", SuperuserEmail).Find(&users, "name LIKE ? OR email LIKE ?", like, like)
		results <- users
	}()
	return <-results
}

func (user *User) Delete() error {
	if user.IsSuperuser() {
		return errActionNotPermitted
	}

	db := database.Conn()
	defer db.Close()

	if err := db.Unscoped().Delete(&user).Error; err != nil {
		return err
	}
	dir := filepath.Join(contentDir, "users", strconv.Itoa(int(user.ID)))

	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		if err := os.RemoveAll(dir); err != nil {
			return err
		}
	}
	return nil
}

func DeleteManyUser(ids []int64) error {
	db := database.Conn()
	defer db.Close()

	for _, id := range ids {
		user, err := UserByID(id)

		if err != nil {
			return err
		}

		if err := user.Delete(); err != nil {
			return err
		}
	}
	return nil
}

func (user *User) Create() (*User, error) {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (user *User) update(update map[string]interface{}) (*User, error) {
	db := database.Conn()
	defer db.Close()

	if err := db.Model(&user).Updates(update).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (user *User) UpdateProfile() error {
	update := make(map[string]interface{})
	update["name"] = user.Name
	update["email"] = user.Email
	update["is_admin"] = user.IsAdmin

	if _, err := user.update(update); err != nil {
		return err
	}
	return nil
}

func (user *User) ResetPassword() error {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return err
	}
	update := make(map[string]interface{})
	update["password"] = string(hashed)

	if _, err := user.update(update); err != nil {
		return err
	}
	return nil
}

func UserByID(id int64) (*User, error) {
	db := database.Conn()
	defer db.Close()

	user := User{}

	if err := db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UserByEmail(email string) (*User, error) {
	db := database.Conn()
	defer db.Close()

	user := User{}

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func UserEmailTaken(email string) (bool, error) {
	user, err := UserByEmail(email)

	if err != nil {
		return false, err
	}

	if user != nil {
		return true, nil
	}
	return false, nil
}

func UserEmailSameAsOld(id int64, email string) (bool, error) {
	user, err := UserByID(id)

	if err != nil {
		return false, err
	}

	if user.Email != email {
		return false, nil
	}
	return true, nil
}

func createSuperUser() error {
	name, _ := env.Conf.String("superuser.name")
	pwd, _ := env.Conf.String("superuser.pwd")
	hashed, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	user := User{
		Name:     name,
		Email:    SuperuserEmail,
		Password: string(hashed),
		IsAdmin:  true,
	}

	if _, err := user.Create(); err != nil {
		return err
	}
	return nil
}

func userSeeder() {
	for i := 0; i < 50; i++ {
		password := "secret"
		hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

		if err != nil {
			panic(err)
		}

		user := User{
			Name:     fake.FullName(),
			Email:    strings.ToLower(fake.EmailAddress()),
			Password: string(hashed),
			IsAdmin:  true,
		}

		if _, err := user.Create(); err != nil {
			panic(err)
		}
	}
}

func GetAuthorizedUser(r *http.Request) (*User, error) {
	ctx := r.Context()
	tokenName, _ := env.Conf.String("pkg.jwtauth.tokenName")
	jwtToken, ok := ctx.Value(tokenName).(*jwt.Token)

	if err, ok := ctx.Value("jwt.err").(error); ok {
		if err != nil {
			return nil, err
		}
	}

	if !ok || jwtToken == nil || !jwtToken.Valid {
		// todo: remove vendor folder on $GOPATH/src/github.com/goware/jwtauth
		err := errors.New("jwt token is either not ok or nil or invalid")
		return nil, err
	}
	claims := jwtToken.Claims
	id := claims["id"].(float64)
	user, err := UserByID(int64(id))

	if err != nil {
		return nil, err
	}
	return user, nil
}

func (user *User) SetPhoto(file multipart.File) error {
	photoPath, _ := env.Conf.String("default.photo.path")
	id := strconv.Itoa(int(user.ID))
	name := filepath.Join(contentDir, "users", id, "photo", filepath.Base(photoPath))

	if err := helper.SaveAsJPEG(file, name); err != nil {
		return err
	}
	db := database.Conn()
	defer db.Close()

	if err := db.Model(&User{}).Where("id = ?", id).Update("has_photo", true).Error; err != nil {
		return err
	}
	return nil
}

func (user *User) GetPhoto() string {
	photoPath, _ := env.Conf.String("default.photo.path")
	return path.Join(contentDir, "users", strconv.Itoa(int(user.ID)), "photo", filepath.Base(photoPath))
}

func (user *User) IsSuperuser() bool {
	if user.Email == SuperuserEmail {
		return true
	}
	return false
}
