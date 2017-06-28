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
	Password   string    `json:"password" bson:"password"`
}

type BasicInfo struct {
	StSub          string     `json:"stSub" bson:"stSub"`
	CityMun        *CityMun   `json:"cityMun" bson:"cityMun"`
	Province       *Province  `json:"province" bson:"province"`
	Barangay       *Barangay  `json:"barangay" bson:"barangay"`
	PlaceOfBirth   string     `json:"placeOfBirth" bson:"placeOfBirth"`
	Religion       *Religion  `json:"religion" bson:"religion"`
	CivilStat      *CivilStat `json:"civilStat" bson:"civilStat"`
	CivilStatOther string     `json:"civilStatOther" bson:"civilStatOther"`
	Sex            *Sex       `json:"sex" bson:"sex"`
	Age            int        `json:"age" bson:"age"`
	Height         float32    `json:"height" bson:"height"`
	Weight         float32    `json:"weight" bson:"weight"`
	LandlineNumber string     `json:"landlineNumber" bson:"landlineNumber"`
	MobileNumber   string     `json:"mobileNumber" bson:"mobileNumber"`
	Email          string     `json:"email" bson:"email"`
}

type Employment struct {
	Stat                     *EmpStat    `json:"stat" bson:"stat"`
	UnEmpStat                *UnEmpStat  `json:"unEmpStatus" bson:"unEmpStat"`
	TeminatedOverseasCountry *Country    `json:"teminatedOverseasCountry" bson:"teminatedOverseasCountry"`
	IsActivelyLookingForWork bool        `json:"isActivelyLookingForWork" bson:"isActivelyLookingForWork"`
	PrefOccs                 []*Position `json:"prefOccs" bson:"prefOccs"`
	PrefLocalLoc             *CityMun    `json:"prefLocalLoc" bson:"prefLocalLoc"`
	PrefOverseasLoc          *Country    `json:"prefOverseasLoc" bson:"prefOverseasLoc"`
	PassportNumber           string      `json:"passportNumber" bson:"passportNumber"`
	PassportNumberExpiryDate time.Time   `json:"passportNumberExpiryDate" bson:"passportNumberExpiryDate"`
}

type Disab struct {
	IsDisabled bool        `json:"isDisabled" bson:"isDisabled"`
	Name       *Disability `json:"name" bson:"name"`
	Other      string      `json:"disab" bson:"Other"`
}

type FormalEdu struct {
	HighestGradeCompleted *EduLevel `json:"highestGradeCompleted" bson:"highestGradeCompleted"`
	CourseDegree          *Course   `json:"courseDegree" bson:"courseDegree"`
	SchoolUniv            *School   `json:"schoolUniv" bson:"schoolUniv"`
	SchoolUnivOther       string    `json:"schoolUnivOther" bson:"schoolUnivOther"`
	YearGrad              time.Time `json:"yearGrad" bson:"yearGrad"`
	LastAttended          time.Time `json:"lastAttended" bson:"lastAttended"`
}

type FormalEduArr struct {
	HighestGradeCompletedId string `json:"formalEduHighestGradeCompletedId"`
	CourseDegreeId          string `json:"formalEduCourseDegreeId"`
	SchoolUnivId            string `json:"formalEduSchoolUnivId"`
	SchoolUnivOther         string `json:"formalEduSchoolUnivOther"`
	YearGrad                string `json:"formalEduYearGrad"`
	LastAttended            string `json:"formalEduLastAttended"`
}

type ProLicense struct {
	Title      *License  `json:"title" bson:"title"`
	ExpiryDate time.Time `json:"expiryDate" bson:"expiryDate"`
}

type ProLicenseArr struct {
	TitleId    string `json:"proLicenseTitleId"`
	ExpiryDate string `json:"proLicenseExpiryDate"`
}

type Elig struct {
	Title     *Eligibility `json:"title" bson:"title"`
	YearTaken time.Time    `json:"yearTaken" bson:"yearTaken"`
}

type EligArr struct {
	TitleId   string `json:"eligTitleId"`
	YearTaken string `json:"eligYearTaken"`
}

type Training struct {
	Name                string `json:"name" bson:"name"`
	SkillsAcquired      string `json:"skillsAcquired" bson:"skillsAcquired"`
	PeriodOfTrainingExp string `json:"periodOfTrainingExp" bson:"periodOfTrainingExp"`
	CertReceived        string `json:"certReceived" bson:"certReceived"`
	IssuingSchoolAgency string `json:"issuingSchoolAgency" bson:"issuingSchoolAgency"`
}

type TrainingArr struct {
	Name                string `json:"trainingNameOfTraining"`
	SkillsAcquired      string `json:"trainingSkillsAcquired"`
	PeriodOfTrainingExp string `json:"trainingPeriodOfTrainingExp"`
	CertReceived        string `json:"trainingCertReceived"`
	IssuingSchoolAgency string `json:"trainingIssuingSchoolAgency"`
}

type Cert struct {
	Title      *Certificate `json:"title" bson:"title"`
	Rating     string       `json:"rating" bson:"rating"`
	IssuedBy   string       `json:"issuedBy" bson:"issuedBy"`
	DateIssued time.Time    `json:"dateIssued" bson:"dateIssued"`
}

type CertArr struct {
	TitleId    string `json:"certTitleId"`
	Rating     string `json:"certRating"`
	IssuedBy   string `json:"certIssuedBy"`
	DateIssued string `json:"certDateIssued"`
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
	NameOfCompanyFirm    string `json:"workExpNameOfCompanyFirm"`
	Address              string `json:"workExpAddress"`
	PositionHeldId       string `json:"workExpPositionHeldId"`
	From                 string `json:"workExpFrom"`
	To                   string `json:"workExpTo"`
	IsRelatedToFormalEdu bool   `json:"workExpIsRelatedToFormalEdu"`
}

type Registrant struct {
	Id              bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	CreatedAt       time.Time     `json:"createdAt" bson:"createdAt"`
	UpdatedAt       time.Time     `json:"updatedAt" bson:"updatedAt"`
	RegisteredAt    time.Time     `json:"registeredAt" bson:"registeredAt"`
	IAccept         bool          `json:"iAccept" bson:"iAccept"`
	PersonalInfo    *PersonalInfo `json:"personalInfo" bson:"personalInfo"`
	BasicInfo       *BasicInfo    `json:"basicInfo" bson:"basicInfo"`
	Employment      *Employment   `json:"employment" bson:"employment"`
	Disab           *Disab        `json:"disab" bson:"disab"`
	Langs           []*Language   `json:"langs" bson:"langs"`
	FormalEdus      []*FormalEdu  `json:"formalEdus" bson:"formalEdus"`
	ProLicenses     []*ProLicense `json:"proLicenses" bson:"proLicenses"`
	Eligs           []*Elig       `json:"eligs" bson:"eligs"`
	Trainings       []*Training   `json:"trainings" bson:"trainings"`
	Certs           []*Cert       `json:"certs" bson:"certs"`
	WorkExps        []*WorkExp    `json:"workExps" bson:"workExps"`
	OtherSkills     []*OtherSkill `json:"otherSkills" bson:"otherSkills"`
	OtherSkillOther string        `json:"otherSkillOther" bson:"otherSkillOther"`
}

type CreateRegistrantForm struct {
	PersonalInfoFamilyName        string                `schema:"personalInfoFamilyName" validate:"required"`
	PersonalInfoGivenName         string                `schema:"personalInfoGivenName" validate:"required"`
	PersonalInfoMiddleName        string                `schema:"personalInfoMiddleName" validate:"required"`
	PersonalInfoBirthdate         string                `schema:"personalInfoBirthdate" validate:"required"`
	PersonalInfoPassword          string                `schema:"personalInfoPassword"`
	PersonalInfoPhotoFile         multipart.File        `schema:"-"`
	PersonalInfoPhotoHeader       *multipart.FileHeader `schema:"-"`
	BasicInfoStSub                string                `schema:"basicInfoStSub" validate:"required"`
	BasicInfoCityMunId            string                `schema:"basicInfoCityMunId" validate:"required"`
	BasicInfoProvinceId           string                `schema:"BasicInfoProvinceId"`
	BasicInfoBarangayId           string                `schema:"basicInfoBarangayId" validate:"required"`
	BasicInfoPlaceOfBirth         string                `schema:"basicInfoPlaceOfBirth" validate:"required"`
	BasicInfoReligionId           string                `schema:"basicInfoReligionId" validate:"required"`
	BasicInfoCivilStatId          string                `schema:"basicInfoCivilStatId" validate:"required"`
	BasicInfoCivilStatOther       string                `schema:"basicInfoCivilStatOther"`
	BasicInfoSexId                string                `schema:"basicInfoSexId" validate:"required"`
	BasicInfoAge                  int                   `schema:"basicInfoAge"`
	BasicInfoHeight               float32               `schema:"basicInfoHeight" validate:"required"`
	BasicInfoWeight               float32               `schema:"basicInfoWeight"`
	BasicInfoLandlineNumber       string                `schema:"basicInfoLandlineNumber"`
	BasicInfoMobileNumber         string                `schema:"basicInfoMobileNumber"`
	BasicInfoEmail                string                `schema:"basicInfoEmail"`
	EmpStatId                     string                `schema:"empStatId" validate:"required"`
	EmpUnEmpStatId                string                `schema:"empUnEmpStatId"`
	EmpTeminatedOverseasCountryId string                `schema:"empTeminatedOverseasCountryId"`
	EmpIsActivelyLookingForWork   bool                  `schema:"empIsActivelyLookingForWork"`
	EmpPrefOccIds                 []string              `schema:"empPrefOccIds" validate:"required"`
	EmpPrefLocalLocId             string                `schema:"empPrefLocalLocId" validate:"required"`
	EmpPrefOverseasLocId          string                `schema:"empPrefOverseasLocId" validate:"required"`
	EmpPassportNumber             string                `schema:"empPassportNumber"`
	EmpPassportNumberExpiryDate   string                `schema:"empPassportNumberExpiryDate"`
	DisabIsDisabled               bool                  `schema:"disabIsDisabled"`
	DisabId                       string                `schema:"disabId"`
	DisabOther                    string                `schema:"disabOther"`
	LangIds                       []string              `schema:"langIds"`
	OtherSkillIds                 []string              `schema:"otherSkillIds"`
	OtherSkillOther               string                `schema:"otherSkillOther"`
	RegisteredAt                  string                `schema:"registeredAt"`
	IAccept                       bool                  `schema:"iAccept"`
	FormalEduJson                 string                `schema:"formalEduJson"`
	ProLicenseJson                string                `schema:"proLicenseJson"`
	EligJson                      string                `schema:"eligJson"`
	TrainingJson                  string                `schema:"trainingJson"`
	CertJson                      string                `schema:"certJson"`
	WorkExpJson                   string                `schema:"workExpJson"`
	Errors                        map[string]string     `schema:"-"`
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

	if err := db.C("registrant").Find(bson.M{"basicInfo.email": email}).One(registrant); err != nil {
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
	hashed, err := bcrypt.GenerateFromPassword([]byte(registrant.PersonalInfo.Password), bcrypt.DefaultCost)

	if err != nil {
		panic(err)
	}
	registrant.PersonalInfo.Password = string(hashed)
	registrant.CreatedAt = time.Now()
	registrant.UpdatedAt = time.Now()

	if err := db.C("registrants").Insert(registrant); err != nil {
		panic(err)
	}
	defer db.Close()
	return registrant
}
