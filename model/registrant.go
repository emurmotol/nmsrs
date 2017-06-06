package model

import "time"

type Registrant struct {
	ID        int64      `json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	RegistInfo   RegistInfo
	RegistEmp    RegistEmp
}

// type CreateRegistrantForm struct {
// 	Name            string                `schema:"name" validate:"required"`
// 	Email           string                `schema:"email" validate:"required,email"`
// 	Password        string                `schema:"password" validate:"required"`
// 	ConfirmPassword string                `schema:"confirm_password" validate:"eqfield=Password"`
// 	IsAdmin         bool                  `schema:"is_admin"`
// 	PhotoFile       multipart.File        `schema:"-"`
// 	PhotoHeader     *multipart.FileHeader `schema:"-"`
// 	Errors          map[string]string     `schema:"-"`
// }

// func (form *CreateRegistrantForm) IsValid() bool {
// 	form.Errors = make(map[string]string)

// 	if errs := helper.ValidateForm(form); len(errs) != 0 {
// 		form.Errors = errs
// 	}

// 	if taken, _ := RegistrantEmailTaken(form.Email); taken {
// 		form.Errors["Email"] = lang.Get("email_taken")
// 	}

// 	if form.PhotoFile != nil {
// 		if err := helper.ValidateImage(form.PhotoHeader); err != nil {
// 			form.Errors["Photo"] = err.Error()
// 		}
// 	}
// 	return len(form.Errors) == 0
// }

// func (registrant Registrant) Search(q string) []Registrant {
// 	db := database.Conn()
// 	defer db.Close()

// 	registrants := []Registrant{}
// 	results := make(chan []Registrant)
// 	like := database.WrapLike(q)

// 	go func() {
// 		db.Find(&registrants, "name LIKE ? OR email LIKE ?", like, like)
// 		results <- registrants
// 	}()
// 	return <-results
// }

// func (registrant *Registrant) Delete() error {
// 	db := database.Conn()
// 	defer db.Close()

// 	if err := db.Unscoped().Delete(&registrant).Error; err != nil {
// 		return err
// 	}
// 	dir := filepath.Join(contentDir, "registrants", strconv.Itoa(int(registrant.ID)))

// 	if _, err := os.Stat(dir); !os.IsNotExist(err) {
// 		if err := os.RemoveAll(dir); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func DeleteManyRegistrant(ids []int64) error {
// 	db := database.Conn()
// 	defer db.Close()

// 	for _, id := range ids {
// 		registrant, err := RegistrantByID(id)

// 		if err != nil {
// 			return err
// 		}

// 		if err := registrant.Delete(); err != nil {
// 			return err
// 		}
// 	}
// 	return nil
// }

// func (registrant *Registrant) Create() (*Registrant, error) {
// 	db := database.Conn()
// 	defer db.Close()

// 	if err := db.Create(&registrant).Error; err != nil {
// 		return nil, err
// 	}
// 	return registrant, nil
// }

// func (registrant *Registrant) update(update map[string]interface{}) (*Registrant, error) {
// 	db := database.Conn()
// 	defer db.Close()

// 	if err := db.Model(&registrant).Updates(update).Error; err != nil {
// 		return nil, err
// 	}
// 	return registrant, nil
// }

// func (registrant *Registrant) UpdateRegistrant() error {
// 	update := make(map[string]interface{})
// 	update["name"] = registrant.Name
// 	update["email"] = registrant.Email
// 	update["is_admin"] = registrant.IsAdmin

// 	if _, err := registrant.update(update); err != nil {
// 		return err
// 	}
// 	return nil
// }

// func RegistrantByID(id int64) (*Registrant, error) {
// 	db := database.Conn()
// 	defer db.Close()

// 	registrant := Registrant{}

// 	if err := db.First(&registrant, id).Error; err != nil {
// 		return nil, err
// 	}
// 	return &registrant, nil
// }

// func RegistrantByEmail(email string) (*Registrant, error) {
// 	db := database.Conn()
// 	defer db.Close()

// 	registrant := Registrant{}

// 	if err := db.Where("email = ?", email).First(&registrant).Error; err != nil {
// 		return nil, err
// 	}
// 	return &registrant, nil
// }

// func RegistrantEmailTaken(email string) (bool, error) {
// 	registrant, err := RegistrantByEmail(email)

// 	if err != nil {
// 		return false, err
// 	}

// 	if registrant != nil {
// 		return true, nil
// 	}
// 	return false, nil
// }

// func RegistrantEmailSameAsOld(id int64, email string) (bool, error) {
// 	registrant, err := RegistrantByID(id)

// 	if err != nil {
// 		return false, err
// 	}

// 	if registrant.Email != email {
// 		return false, nil
// 	}
// 	return true, nil
// }

// func RegistrantSeeder() {
// 	for i := 0; i < 50; i++ {
// 		password := "secret"
// 		hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

// 		if err != nil {
// 			panic(err)
// 		}

// 		registrant := Registrant{
// 			Name:     fake.FullName(),
// 			Email:    strings.ToLower(fake.EmailAddress()),
// 			Password: string(hashed),
// 			IsAdmin:  true,
// 		}
// 		registrant.Create()
// 	}
// }

// func (registrant *Registrant) SetPhoto(file multipart.File) error {
// 	photoPath, _ := env.Conf.String("default.photo.path")
// 	id := strconv.Itoa(int(registrant.ID))
// 	name := filepath.Join(contentDir, "registrants", id, "photo", filepath.Base(photoPath))

// 	if err := helper.SaveAsJPEG(file, name); err != nil {
// 		return err
// 	}
// 	db := database.Conn()
// 	defer db.Close()

// 	if err := db.Model(&Registrant{}).Where("id = ?", id).Update("has_photo", true).Error; err != nil {
// 		return err
// 	}
// 	return nil
// }

// func (registrant *Registrant) GetPhoto() string {
// 	photoPath, _ := env.Conf.String("default.photo.path")
// 	return path.Join(contentDir, "registrants", strconv.Itoa(int(registrant.ID)), "photo", filepath.Base(photoPath))
// }
