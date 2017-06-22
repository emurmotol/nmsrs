package model

// import (
// 	"mime/multipart"
// 	"time"

// 	"github.com/emurmotol/nmsrs/db"
// 	"github.com/emurmotol/nmsrs/helper"
// 	"github.com/emurmotol/nmsrs/lang"
// )

// type Registrant struct {
// 	Id           uint64     `json:"id"`
// 	CreatedAt    time.Time  `json:"created_at"`
// 	UpdatedAt    time.Time  `json:"updated_at"`
// 	DeletedAt    *time.Time `json:"deleted_at"`
// 	RegisteredAt string     `json:"registered_at"`
// 	IAccept      bool       `json:"i_accept"`
// 	RegistInfo   RegistInfo
// 	RegistEmp    RegistEmp
// }

// type CreateRegistrantForm struct {
// 	FamilyName        string                `schema:"family_name" validate:"required"`
// 	GivenName         string                `schema:"given_name" validate:"required"`
// 	MiddleName        string                `schema:"middle_name" validate:"required"`
// 	Birthdate         string                `schema:"birthdate" validate:"required"`
// 	Password          string                `schema:"password"`
// 	PhotoFile         multipart.File        `schema:"-"`
// 	PhotoHeader       *multipart.FileHeader `schema:"-"`
// 	StSub             string                `schema:"st_sub" validate:"required"`
// 	CityMunId         uint                  `schema:"city_mun_id" validate:"required"`
// 	ProvId            uint                  `schema:"prov_id"`
// 	BrgyId            uint                  `schema:"brgy_id" validate:"required"`
// 	PlaceOfBirth      string                `schema:"place_of_birth" validate:"required"`
// 	ReligionId        uint                  `schema:"religion_id" validate:"required"`
// 	CivilStatId       uint                  `schema:"civil_stat_id" validate:"required"`
// 	CivilStatOther    string                `schema:"civil_stat_other"`
// 	SexId             uint                  `schema:"sex_id" validate:"required"`
// 	Age               int                   `schema:"age"`
// 	Height            float32               `schema:"height" validate:"required"`
// 	Weight            float32               `schema:"weight"`
// 	LandlineNo        string                `schema:"landline_no"`
// 	MobileNo          string                `schema:"mobile_no"`
// 	Email             string                `schema:"email"`
// 	EmpStatId         uint                  `schema:"emp_stat_id" validate:"required"`
// 	UnEmpStatId       uint                  `schema:"un_emp_stat_id"`
// 	TocId             uint                  `schema:"toc_id"`
// 	Alfw              bool                  `schema:"alfw"`
// 	PrefOccIds        []int                 `schema:"pref_occ_ids" validate:"required"`
// 	PrefLocalLocId    uint                  `schema:"pref_local_loc_id" validate:"required"`
// 	PrefOverseasLocId uint                  `schema:"pref_overseas_loc_id" validate:"required"`
// 	PassportNo        string                `schema:"passport_no"`
// 	Pned              string                `schema:"pned"`
// 	Disabled          bool                  `schema:"disabled"`
// 	DisabilityId      uint                  `schema:"disability_id"`
// 	DisabilityOther   uint                  `schema:"disability_other"`
// 	LanguageIds       []int                 `schema:"language_ids"`
// 	SkillIds          []int                 `schema:"skill_ids"`
// 	RegisteredAt      string                `schema:"registered_at"`
// 	IAccept           bool                  `schema:"i_accept"`
// 	Errors            map[string]string     `schema:"-"`
// }

// func (form *CreateRegistrantForm) IsValid() bool {
// 	form.Errors = make(map[string]string)

// 	if errs := helper.ValidateForm(form); len(errs) != 0 {
// 		form.Errors = errs
// 	}

// 	if taken := RegistrantEmailTaken(form.Email); taken {
// 		form.Errors["Email"] = lang.Get("email_taken")
// 	}

// 	if form.PhotoFile != nil {
// 		if err := helper.ValidateImage(form.PhotoHeader); err != nil {
// 			form.Errors["Photo"] = err.Error()
// 		}
// 	}
// 	return len(form.Errors) == 0
// }

// func RegistrantById(id uint64) *Registrant {
// 	db := database.Con()
// 	defer db.Close()
// 	registrant := new(Registrant)

// 	if notFound := db.First(registrant, id).RecordNotFound(); notFound {
// 		return nil
// 	}
// 	return registrant
// }

// func RegistrantByEmail(email string) *Registrant {
// 	db := database.Con()
// 	defer db.Close()
// 	registInfo := RegistInfo{}

// 	if notFound := db.Where("email = ?", email).First(&registInfo).RecordNotFound(); notFound {
// 		return nil
// 	}
// 	return RegistrantById(registInfo.RegistrantId)
// }

// func RegistrantEmailTaken(email string) bool {
// 	registrant := RegistrantByEmail(email)

// 	if registrant != nil {
// 		return true
// 	}
// 	return false
// }

// func (registrant *Registrant) Create() *Registrant {
// 	db := database.Con()
// 	defer db.Close()

// 	if err := db.Create(&registrant).Error; err != nil {
// 		panic(err)
// 	}
// 	return registrant
// }
