package model

import (
	"mime/multipart"
	"time"
)

type Registrant struct {
	ID         uint64     `json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `json:"deleted_at"`
	RegistInfo RegistInfo
	RegistEmp  RegistEmp
}

type CreateRegistrantForm struct {
	FamilyName        string                `schema:"family_name" validate:"required"`
	GivenName         string                `schema:"given_name" validate:"required"`
	MiddleName        string                `schema:"middle_name" validate:"required"`
	Birthdate         time.Time             `schema:"birthdate" validate:"required"`
	Password          string                `schema:"password"`
	PhotoFile         multipart.File        `schema:"-"`
	PhotoHeader       *multipart.FileHeader `schema:"-"`
	StSub             string                `schema:"st_sub" validate:"required"`
	CityMunID         uint                  `schema:"city_mun_id" validate:"required"`
	ProvID            uint                  `schema:"prov_id"`
	BrgyID            uint                  `schema:"brgy_id" validate:"required"`
	PlaceOfBirth      string                `schema:"place_of_birth" validate:"required"`
	ReligionID        uint                  `schema:"religion_id" validate:"required"`
	CivilStatID       uint                  `schema:"civil_stat_id" validate:"required"`
	CivilStatOther    string                `schema:"civil_stat_other" validate:"required"`
	SexID             uint                  `schema:"sex_id" validate:"required"`
	Age               int                   `schema:"age"`
	Height            float32               `schema:"height" validate:"required"`
	Weight            float32               `schema:"weight"`
	LandlineNo        string                `schema:"landline_no"`
	MobileNo          string                `schema:"mobile_no"`
	Email             string                `schema:"email" validate:"email"`
	EmpStatID         uint                  `schema:"emp_stat_id" validate:"required"`
	UnEmpStatID       uint                  `schema:"un_emp_stat_id"`
	TocID             uint                  `schema:"toc_id"`
	Alfw              bool                  `schema:"alfw"`
	PrefOccIDs        []int                 `schema:"pref_occ_ids" validate:"required"`
	PrefLocalLocID    uint                  `schema:"pref_local_loc_id" validate:"required"`
	PrefOverseasLocID uint                  `schema:"pref_overseas_loc_id" validate:"required"`
	PassportNo        string                `schema:"passport_no"`
	Pned              time.Time             `schema:"pned"`
	Disabled          bool                  `schema:"disabled"`
	DisabilityID      uint                  `schema:"disability_id"`
	DisabilityOther   uint                  `schema:"disability_other"`
	LanguageIDs       []int                 `schema:"language_ids"`
	Errors            map[string]string     `schema:"-"`
}

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

// func DeleteManyRegistrant(ids []uint64) error {
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

// func RegistrantByID(id uint64) (*Registrant, error) {
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

// func RegistrantEmailSameAsOld(id uint64, email string) (bool, error) {
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
