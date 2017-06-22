package model

import (
	"errors"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"

	"mime/multipart"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/emurmotol/nmsrs/db"
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
	Id        bson.ObjectId `json:"id" bson:"_id"`
	CreatedAt time.Time     `json:"created_at" bson:"createdAt"`
	UpdatedAt time.Time     `json:"updated_at" bson:"updatedAt"`
	Name      string        `json:"name" bson:"name"`
	Email     string        `json:"email" bson:"email"`
	Password  string        `json:"password" bson:"password"`
	IsAdmin   bool          `json:"is_admin" bson:"isAdmin"`
	HasPhoto  bool          `json:"has_photo" bson:"hasPhoto"`
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

	if taken := UserEmailTaken(form.Email); !taken {
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

func (createUserForm *CreateUserForm) IsValid() bool {
	createUserForm.Errors = make(map[string]string)

	if errs := helper.ValidateForm(createUserForm); len(errs) != 0 {
		createUserForm.Errors = errs
	}

	if taken := UserEmailTaken(createUserForm.Email); taken {
		createUserForm.Errors["Email"] = lang.Get("email_taken")
	}

	if createUserForm.PhotoFile != nil {
		if err := helper.ValidateImage(createUserForm.PhotoHeader); err != nil {
			createUserForm.Errors["Photo"] = err.Error()
		}
	}
	return len(createUserForm.Errors) == 0
}

type EditProfileForm struct {
	IdHex       string                `schema:"-"`
	Name        string                `schema:"name" validate:"required"`
	Email       string                `schema:"email" validate:"required,email"`
	IsAdmin     bool                  `schema:"is_admin"`
	PhotoFile   multipart.File        `schema:"-"`
	PhotoHeader *multipart.FileHeader `schema:"-"`
	Errors      map[string]string     `schema:"-"`
}

func (editProfileForm *EditProfileForm) IsValid() bool {
	editProfileForm.Errors = make(map[string]string)

	if errs := helper.ValidateForm(editProfileForm); len(errs) != 0 {
		editProfileForm.Errors = errs
	}

	if same := UserEmailSameAsOld(bson.ObjectIdHex(editProfileForm.IdHex), editProfileForm.Email); !same {
		if taken := UserEmailTaken(editProfileForm.Email); taken {
			editProfileForm.Errors["Email"] = lang.Get("email_taken")
		}
	}

	if editProfileForm.PhotoFile != nil {
		if err := helper.ValidateImage(editProfileForm.PhotoHeader); err != nil {
			editProfileForm.Errors["Photo"] = err.Error()
		}
	}
	return len(editProfileForm.Errors) == 0
}

type PasswordResetForm struct {
	NewPassword     string            `schema:"new_password" validate:"required"`
	ConfirmPassword string            `schema:"confirm_password" validate:"eqfield=NewPassword"`
	Errors          map[string]string `schema:"-"`
}

func (passwordResetForm *PasswordResetForm) IsValid() bool {
	passwordResetForm.Errors = make(map[string]string)

	if errs := helper.ValidateForm(passwordResetForm); len(errs) != 0 {
		passwordResetForm.Errors = errs
	}
	return len(passwordResetForm.Errors) == 0
}

func (user User) Search(q string) []User {
	users := []User{}
	r := make(chan []User)
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{
		"$or": []bson.M{
			bson.M{"name": regex},
			bson.M{"email": regex},
		},
	}

	go func() {
		db.C("users").Find(query).All(&users)
		defer db.Close()
		r <- users
	}()

	users = <-r
	close(r)
	return users
}

func (user *User) Delete() {
	if user.IsSuperuser() {
		panic(errActionNotPermitted)
	}
	db.C("users").RemoveId(user.Id)
	defer db.Close()
	dir := filepath.Join(contentDir, "users", user.Id.Hex())

	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		if err := os.RemoveAll(dir); err != nil {
			panic(err)
		}
	}
}

func DeleteManyUser(hexIds []string) {
	for _, hexId := range hexIds {
		user := UserById(bson.ObjectIdHex(hexId))
		user.Delete()
	}
}

func (user *User) Create() *User {
	user.Id = bson.NewObjectId()
	user.Email = strings.ToLower(user.Email)
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}
	user.Password = string(hashed)
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	db.C("users").Insert(user)
	defer db.Close()
	return user
}

func (user *User) UpdateProfile() {
	if user.IsSuperuser() {
		panic(errActionNotPermitted)
	}
	user.Email = strings.ToLower(user.Email)
	user.UpdatedAt = time.Now()
	db.C("users").UpdateId(user.Id, bson.M{"$set": user})
	defer db.Close()
}

func (user *User) ResetPassword() {
	hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}
	user.Password = string(hashed)
	db.C("users").UpdateId(user.Id, bson.M{"$set": user})
	defer db.Close()
}

func UserById(id bson.ObjectId) *User {
	user := new(User)
	db.C("users").FindId(id).One(&user)
	defer db.Close()
	return user
}

func UserByEmail(email string) *User {
	user := new(User)
	db.C("users").Find(bson.M{"email": email}).One(&user)
	defer db.Close()
	return user
}

func UserEmailTaken(email string) bool {
	user := UserByEmail(email)

	if user != nil {
		return true
	}
	return false
}

func UserEmailSameAsOld(id bson.ObjectId, email string) bool {
	user := UserById(id)

	if user.Email != email {
		return false
	}
	return true
}

func createSuperUser() {
	name, _ := env.Conf.String("superuser.name")
	pwd, _ := env.Conf.String("superuser.pwd")
	hashed, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}

	user := User{
		Name:     name,
		Email:    SuperUserEmail,
		Password: string(hashed),
		IsAdmin:  true,
	}
	user.Create()
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
		user.Create()
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
	id := claims["userId"].(string)
	return UserById(bson.ObjectIdHex(id)), nil
}

func (user *User) SetPhoto(file multipart.File) error {
	photoPath, _ := env.Conf.String("default.photo.path")
	name := filepath.Join(contentDir, "users", user.Id.Hex(), "photo", filepath.Base(photoPath))

	if err := helper.SaveAsJPEG(file, name); err != nil {
		return err
	}
	user.HasPhoto = true

	if err := db.C("users").UpdateId(user.Id, bson.M{"$set": user}); err != nil {
		return err
	}
	defer db.Close()
	return nil
}

func (user *User) GetPhoto() string {
	photoPath, _ := env.Conf.String("default.photo.path")
	return path.Join(contentDir, "users", user.Id.Hex(), "photo", filepath.Base(photoPath))
}

func (user *User) IsSuperuser() bool {
	if user.Email == SuperUserEmail {
		return true
	}
	return false
}
