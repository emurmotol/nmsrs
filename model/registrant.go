package model

import (
	"mime/multipart"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/helper"
	"github.com/emurmotol/nmsrs/lang"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type PersonalInfo struct {
	HasPhoto   bool      `json:"hasPhoto" bson:"hasPhoto"`
	FamilyName string    `json:"familyName" bson:"familyName"`
	GivenName  string    `json:"givenName" bson:"givenName"`
	MiddleName string    `json:"middleName" bson:"middleName"`
	Birthdate  time.Time `json:"birthdate" bson:"birthdate"`
	Password   string    `json:"password,omitempty" bson:"password,omitempty"`
}

type Height struct {
	Feet   float32 `json:"feet,omitempty" bson:"feet,omitempty"`
	Inches float32 `json:"inches,omitempty" bson:"inches,omitempty"`
}

type BasicInfo struct {
	StSub          string     `json:"stSub,omitempty" bson:"stSub,omitempty"`
	CityMun        *CityMun   `json:"cityMun,omitempty" bson:"cityMun,omitempty"`
	Province       *Province  `json:"province" bson:"province"`
	Barangay       *Barangay  `json:"barangay" bson:"barangay"`
	PlaceOfBirth   string     `json:"placeOfBirth,omitempty" bson:"placeOfBirth,omitempty"`
	Religion       *Religion  `json:"religion,omitempty" bson:"religion,omitempty"`
	CivilStat      *CivilStat `json:"civilStat" bson:"civilStat"`
	CivilStatOther string     `json:"civilStatOther,omitempty" bson:"civilStatOther,omitempty"`
	Sex            *Sex       `json:"sex" bson:"sex"`
	Age            int        `json:"age,omitempty" bson:"age,omitempty"`
	Height         *Height    `json:"height,omitempty" bson:"height,omitempty"`
	Weight         float32    `json:"weight,omitempty" bson:"weight,omitempty"`
	LandlineNumber string     `json:"landlineNumber,omitempty" bson:"landlineNumber,omitempty"`
	MobileNumber   string     `json:"mobileNumber,omitempty" bson:"mobileNumber,omitempty"`
	Email          string     `json:"email,omitempty" bson:"email,omitempty"`
}

type Employment struct {
	Stat                     *EmpStat    `json:"stat,omitempty" bson:"stat,omitempty"`
	UnEmpStat                *UnEmpStat  `json:"unEmpStatus,omitempty" bson:"unEmpStat,omitempty"`
	TeminatedOverseasCountry *Country    `json:"teminatedOverseasCountry,omitempty" bson:"teminatedOverseasCountry,omitempty"`
	IsActivelyLookingForWork bool        `json:"isActivelyLookingForWork" bson:"isActivelyLookingForWork"`
	PrefOccs                 []*Position `json:"prefOccs,omitempty" bson:"prefOccs,omitempty"`
	PrefLocalLoc             *CityMun    `json:"prefLocalLoc,omitempty" bson:"prefLocalLoc,omitempty"`
	PrefOverseasLoc          *Country    `json:"prefOverseasLoc,omitempty" bson:"prefOverseasLoc,omitempty"`
	PassportNumber           string      `json:"passportNumber,omitempty" bson:"passportNumber,omitempty"`
	PassportNumberExpiryDate time.Time   `json:"passportNumberExpiryDate,omitempty" bson:"passportNumberExpiryDate,omitempty"`
}

type Disab struct {
	IsDisabled bool        `json:"isDisabled" bson:"isDisabled"`
	Name       *Disability `json:"name" bson:"name"`
	Other      string      `json:"other,omitempty" bson:"other,omitempty"`
}

type FormalEdu struct {
	HighestGradeCompleted *EduLevel `json:"highestGradeCompleted" bson:"highestGradeCompleted"`
	CourseDegree          *Course   `json:"courseDegree" bson:"courseDegree"`
	SchoolUniv            *School   `json:"schoolUniv" bson:"schoolUniv"`
	SchoolUnivOther       string    `json:"schoolUnivOther,omitempty" bson:"schoolUnivOther,omitempty"`
	YearGrad              time.Time `json:"yearGrad" bson:"yearGrad"`
	LastAttended          time.Time `json:"lastAttended" bson:"lastAttended"`
}

type FormalEduArr struct {
	HighestGradeCompletedHexId string `json:"highestGradeCompletedHexId"`
	CourseDegreeHexId          string `json:"courseDegreeHexId"`
	SchoolUnivHexId            string `json:"schoolUnivHexId"`
	SchoolUnivOther            string `json:"schoolUnivOther"`
	YearGrad                   int    `json:"yearGrad"`
	LastAttended               string `json:"lastAttended"`
}

type ProLicense struct {
	Title      *License  `json:"title" bson:"title"`
	ExpiryDate time.Time `json:"expiryDate" bson:"expiryDate"`
}

type ProLicenseArr struct {
	TitleHexId string `json:"titleHexId"`
	ExpiryDate string `json:"expiryDate"`
}

type Elig struct {
	Title     *Eligibility `json:"title" bson:"title"`
	YearTaken time.Time    `json:"yearTaken" bson:"yearTaken"`
}

type EligArr struct {
	TitleHexId string `json:"titleHexId"`
	YearTaken  string `json:"yearTaken"`
}

type Training struct {
	Name                string `json:"name" bson:"name"`
	SkillsAcquired      string `json:"skillsAcquired,omitempty" bson:"skillsAcquired,omitempty"`
	PeriodOfTrainingExp string `json:"periodOfTrainingExp" bson:"periodOfTrainingExp"`
	CertReceived        string `json:"certReceived,omitempty" bson:"certReceived,omitempty"`
	IssuingSchoolAgency string `json:"issuingSchoolAgency" bson:"issuingSchoolAgency"`
}

type TrainingArr struct {
	Name                string `json:"nameOfTraining"`
	SkillsAcquired      string `json:"skillsAcquired"`
	PeriodOfTrainingExp string `json:"periodOfTrainingExp"`
	CertReceived        string `json:"certReceived"`
	IssuingSchoolAgency string `json:"issuingSchoolAgency"`
}

type Cert struct {
	Title      *Certificate `json:"title" bson:"title"`
	Rating     string       `json:"rating,omitempty" bson:"rating,omitempty"`
	IssuedBy   string       `json:"issuedBy" bson:"issuedBy"`
	DateIssued time.Time    `json:"dateIssued" bson:"dateIssued"`
}

type CertArr struct {
	TitleHexId string `json:"titleHexId"`
	Rating     string `json:"rating"`
	IssuedBy   string `json:"issuedBy"`
	DateIssued string `json:"dateIssued"`
}

type WorkExp struct {
	NameOfCompanyFirm    string    `json:"nameOfCompanyFirm" bson:"nameOfCompanyFirm"`
	Address              string    `json:"address" bson:"address"`
	PositionHeld         *Position `json:"positionHeld" bson:"positionHeld"`
	From                 time.Time `json:"from" bson:"from"`
	To                   time.Time `json:"to" bson:"to"`
	IsRelatedToFormalEdu bool      `json:"isRelatedToFormalEdu" bson:"isRelatedToFormalEdu"`
}

type WorkExpArr struct {
	NameOfCompanyFirm    string `json:"nameOfCompanyFirm"`
	Address              string `json:"address"`
	PositionHeldHexId    string `json:"positionHeldHexId"`
	From                 string `json:"from"`
	To                   string `json:"to"`
	IsRelatedToFormalEdu bool   `json:"isRelatedToFormalEdu"`
}

type Registrant struct {
	Id              bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt       time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time     `json:"updatedAt" bson:"updatedAt"`
	RegisteredAt    time.Time     `json:"registeredAt" bson:"registeredAt"`
	IAccept         bool          `json:"iAccept" bson:"iAccept"`
	PersonalInfo    *PersonalInfo `json:"personalInfo,omitempty" bson:"personalInfo,omitempty"`
	BasicInfo       *BasicInfo    `json:"basicInfo,omitempty" bson:"basicInfo,omitempty"`
	Employment      *Employment   `json:"employment,omitempty" bson:"employment,omitempty"`
	Disab           *Disab        `json:"disab,omitempty" bson:"disab,omitempty"`
	Langs           []*Language   `json:"langs,omitempty" bson:"langs,omitempty"`
	FormalEdus      []*FormalEdu  `json:"formalEdus,omitempty" bson:"formalEdus,omitempty"`
	ProLicenses     []*ProLicense `json:"proLicenses,omitempty" bson:"proLicenses,omitempty"`
	Eligs           []*Elig       `json:"eligs,omitempty" bson:"eligs,omitempty"`
	Trainings       []*Training   `json:"trainings,omitempty" bson:"trainings,omitempty"`
	Certs           []*Cert       `json:"certs,omitempty" bson:"certs,omitempty"`
	WorkExps        []*WorkExp    `json:"workExps,omitempty" bson:"workExps,omitempty"`
	OtherSkills     []*OtherSkill `json:"otherSkills,omitempty" bson:"otherSkills,omitempty"`
	OtherSkillOther string        `json:"otherSkillOther,omitempty" bson:"otherSkillOther,omitempty"`
}

type CreateRegistrantForm struct {
	PersonalInfoFamilyName           string                `schema:"personalInfoFamilyName" validate:"required"`
	PersonalInfoGivenName            string                `schema:"personalInfoGivenName" validate:"required"`
	PersonalInfoMiddleName           string                `schema:"personalInfoMiddleName" validate:"required"`
	PersonalInfoBirthdate            string                `schema:"personalInfoBirthdate" validate:"required"`
	PersonalInfoPassword             string                `schema:"personalInfoPassword"`
	PersonalInfoPhotoFile            multipart.File        `schema:"-"`
	PersonalInfoPhotoHeader          *multipart.FileHeader `schema:"-"`
	BasicInfoStSub                   string                `schema:"basicInfoStSub"`
	BasicInfoCityMunHexId            string                `schema:"basicInfoCityMunHexId" validate:"required"`
	BasicInfoProvinceHexId           string                `schema:"basicInfoProvinceHexId"`
	BasicInfoBarangayHexId           string                `schema:"basicInfoBarangayHexId" validate:"required"`
	BasicInfoPlaceOfBirth            string                `schema:"basicInfoPlaceOfBirth"`
	BasicInfoReligionHexId           string                `schema:"basicInfoReligionHexId"`
	BasicInfoCivilStatHexId          string                `schema:"basicInfoCivilStatHexId" validate:"required"`
	BasicInfoCivilStatOther          string                `schema:"basicInfoCivilStatOther"`
	BasicInfoSexHexId                string                `schema:"basicInfoSexHexId" validate:"required"`
	BasicInfoAge                     int                   `schema:"basicInfoAge"`
	BasicInfoHeightInFeet            float32               `schema:"basicInfoHeightInFeet"`
	BasicInfoHeightInInches          float32               `schema:"basicInfoHeightInInches"`
	BasicInfoWeight                  float32               `schema:"basicInfoWeight"`
	BasicInfoLandlineNumber          string                `schema:"basicInfoLandlineNumber"`
	BasicInfoMobileNumber            string                `schema:"basicInfoMobileNumber"`
	BasicInfoEmail                   string                `schema:"basicInfoEmail"`
	EmpStatHexId                     string                `schema:"empStatHexId"`
	EmpUnEmpStatHexId                string                `schema:"empUnEmpStatHexId"`
	EmpTeminatedOverseasCountryHexId string                `schema:"empTeminatedOverseasCountryHexId"`
	EmpIsActivelyLookingForWork      bool                  `schema:"empIsActivelyLookingForWork"`
	EmpPrefOccHexIds                 []string              `schema:"empPrefOccHexIds"`
	EmpPrefLocalLocHexId             string                `schema:"empPrefLocalLocHexId"`
	EmpPrefOverseasLocHexId          string                `schema:"empPrefOverseasLocHexId"`
	EmpPassportNumber                string                `schema:"empPassportNumber"`
	EmpPassportNumberExpiryDate      string                `schema:"empPassportNumberExpiryDate"`
	DisabIsDisabled                  bool                  `schema:"disabIsDisabled"`
	DisabHexId                       string                `schema:"disabHexId"`
	DisabOther                       string                `schema:"disabOther"`
	LangHexIds                       []string              `schema:"langHexIds"`
	OtherSkillHexIds                 []string              `schema:"otherSkillHexIds"`
	OtherSkillOther                  string                `schema:"otherSkillOther"`
	RegisteredAt                     string                `schema:"registeredAt"`
	IAccept                          bool                  `schema:"iAccept"`
	FormalEduJson                    string                `schema:"formalEduJson"`
	ProLicenseJson                   string                `schema:"proLicenseJson"`
	EligJson                         string                `schema:"eligJson"`
	TrainingJson                     string                `schema:"trainingJson"`
	CertJson                         string                `schema:"certJson"`
	WorkExpJson                      string                `schema:"workExpJson"`
	Errors                           map[string]string     `schema:"-"`
}

func (createRegistrantForm *CreateRegistrantForm) IsValid() bool {
	createRegistrantForm.Errors = make(map[string]string)

	if errs := helper.ValidateForm(createRegistrantForm); len(errs) != 0 {
		createRegistrantForm.Errors = errs
	}

	if taken := RegistrantEmailTaken(createRegistrantForm.BasicInfoEmail); taken {
		createRegistrantForm.Errors["BasicInfoEmail"] = lang.Get("emailTaken")
	}

	if createRegistrantForm.PersonalInfoPhotoFile != nil {
		if err := helper.ValidateImage(createRegistrantForm.PersonalInfoPhotoHeader); err != nil {
			createRegistrantForm.Errors["Photo"] = err.Error()
		}
	}
	return len(createRegistrantForm.Errors) == 0
}

func RegistrantByEmail(email string) *Registrant {
	registrant := new(Registrant)
	query := bson.M{
		"$and": []bson.M{
			bson.M{
				"basicInfo.email": bson.M{
					"$exists": true,
				},
			},
			bson.M{
				"basicInfo.email": email,
			},
		},
	}

	if err := db.C("registrants").Find(query).One(registrant); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return registrant
}

func RegistrantEmailTaken(email string) bool {
	registrant := RegistrantByEmail(email)

	if registrant != nil {
		return true
	}
	return false
}

func (registrant *Registrant) Create() *Registrant {
	registrant.Id = bson.NewObjectId()

	if registrant.PersonalInfo.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(registrant.PersonalInfo.Password), bcrypt.DefaultCost)

		if err != nil {
			panic(err)
		}
		registrant.PersonalInfo.Password = string(hashed)
	}
	registrant.CreatedAt = time.Now()
	registrant.UpdatedAt = time.Now()

	if err := db.C("registrants").Insert(registrant); err != nil {
		panic(err)
	}
	defer db.Close()
	return registrant
}
