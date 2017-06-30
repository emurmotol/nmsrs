package controller

import (
	"encoding/json"
	"html/template"
	"net/http"

	"gopkg.in/mgo.v2/bson"

	"strconv"

	"strings"

	"github.com/emurmotol/nmsrs/helper"
	"github.com/emurmotol/nmsrs/lang"
	"github.com/emurmotol/nmsrs/model"
	"github.com/unrolled/render"
)

// func GetRegistrants(w http.ResponseWriter, r *http.Request) {
// 	db := database.Con()
// 	defer db.Close()

// 	query := db.Model(&model.Registrant{})
// 	query.Count(&count)
// 	page, err := strconv.Atoi(r.URL.Query().Get("page"))

// 	if err != nil {
// 		page = 1
// 	}

// 	pagination := &helper.Paginator{
// 		Page:     page,
// 		Limit:    limit,
// 		Count:    count,
// 		Interval: interval,
// 		QueryUrl: r.URL.Query(),
// 	}

// 	if page > pagination.PageCount() {
// 		pagination.Page = 1
// 	}
// 	registrants := []model.Registrant{}
// 	query.Offset(pagination.Offset()).Limit(limit).Find(&registrants)

// 	data := make(map[string]interface{})
// 	data["title"] = "Registrants"
// 	data["auth"] = model.Auth(r)
// 	data["registrants"] = registrants
// 	data["q"] = r.URL.Query().Get("q")
// 	data["pagination"] = helper.Pager{
// 		Markup:     template.HTML(pagination.String()),
// 		Count:      pagination.Count,
// 		StartIndex: pagination.StartIndex(),
// 		EndIndex:   pagination.EndIndex(),
// 	}
// 	flashAlert := helper.GetFlash(w, r, "alert")

// 	if flashAlert != nil {
// 		alert := flashAlert.(helper.Alert)
// 		data["alert"] = template.HTML(alert.String())
// 	}
// 	rd.HTML(w, http.StatusOK, "registrant/index", data)
// }

func CreateRegistrant(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	flashAlert := helper.GetFlash(w, r, "alert")

	if flashAlert != nil {
		alert := flashAlert.(helper.Alert)
		data["alert"] = template.HTML(alert.String())
	}
	createRegistrantForm := helper.GetFlash(w, r, "createRegistrantForm")

	if createRegistrantForm != nil {
		data["createRegistrantForm"] = createRegistrantForm.(model.CreateRegistrantForm)
	}
	data["civilStats"] = model.CivilStats()
	data["sexes"] = model.Sexes()
	data["empStats"] = model.EmpStats()
	data["disabilities"] = model.Disabilities()
	data["title"] = "Create Registrant"
	data["auth"] = model.Auth(r)
	rd.HTML(w, http.StatusOK, "registrant/create", data, render.HTMLOptions{Layout: "layouts/wizard"})
}

func StoreRegistrant(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(0); err != nil {
		panic(err)
	}
	photoFile, photoHeader, err := r.FormFile("personalInfoPhoto")

	if err != nil {
		if err != http.ErrMissingFile {
			panic(err)
		}
	}
	delete(r.PostForm, "personalInfoPhoto")
	createRegistrantForm := model.CreateRegistrantForm{}

	if err := decoder.Decode(&createRegistrantForm, r.PostForm); err != nil {
		panic(err)
	}
	createRegistrantForm.PersonalInfoPhotoFile = photoFile
	createRegistrantForm.PersonalInfoPhotoHeader = photoHeader

	if !createRegistrantForm.IsValid() {
		helper.SetFlash(w, r, "createRegistrantForm", createRegistrantForm)
		CreateRegistrant(w, r)
		return
	}
	hasPhoto := false

	if createRegistrantForm.PersonalInfoPhotoFile != nil {
		hasPhoto = true
	}

	registrant := model.Registrant{
		RegisteredAt: helper.ShortDate(createRegistrantForm.RegisteredAt),
		IAccept:      createRegistrantForm.IAccept,
		PersonalInfo: &model.PersonalInfo{
			HasPhoto:   hasPhoto,
			FamilyName: strings.ToUpper(createRegistrantForm.PersonalInfoFamilyName),
			GivenName:  strings.ToUpper(createRegistrantForm.PersonalInfoGivenName),
			MiddleName: strings.ToUpper(createRegistrantForm.PersonalInfoMiddleName),
			Birthdate:  helper.ShortDate(createRegistrantForm.PersonalInfoBirthdate),
			Password:   strings.ToUpper(createRegistrantForm.PersonalInfoPassword),
		},
		BasicInfo: &model.BasicInfo{
			StSub:          strings.ToUpper(createRegistrantForm.BasicInfoStSub),
			Province:       model.ProvinceById(bson.ObjectIdHex(createRegistrantForm.BasicInfoProvinceId)),
			Barangay:       model.BarangayById(bson.ObjectIdHex(createRegistrantForm.BasicInfoBarangayId)),
			PlaceOfBirth:   strings.ToUpper(createRegistrantForm.BasicInfoPlaceOfBirth),
			CivilStat:      model.CivilStatById(bson.ObjectIdHex(createRegistrantForm.BasicInfoCivilStatId)),
			CivilStatOther: strings.ToUpper(createRegistrantForm.BasicInfoCivilStatOther),
			Sex:            model.SexById(bson.ObjectIdHex(createRegistrantForm.BasicInfoSexId)),
			LandlineNumber: createRegistrantForm.BasicInfoLandlineNumber,
			MobileNumber:   createRegistrantForm.BasicInfoMobileNumber,
			Email:          strings.ToLower(createRegistrantForm.BasicInfoEmail),
		},
		Employment: &model.Employment{
			IsActivelyLookingForWork: createRegistrantForm.EmpIsActivelyLookingForWork,
			PassportNumber:           createRegistrantForm.EmpPassportNumber,
		},
		OtherSkillOther: strings.ToUpper(createRegistrantForm.OtherSkillOther),
	}

	if bson.IsObjectIdHex(createRegistrantForm.EmpStatId) {
		registrant.Employment.Stat = model.EmpStatById(bson.ObjectIdHex(createRegistrantForm.EmpStatId))
	}

	if bson.IsObjectIdHex(createRegistrantForm.EmpPrefLocalLocId) {
		registrant.Employment.PrefLocalLoc = model.CityMunById(bson.ObjectIdHex(createRegistrantForm.EmpPrefLocalLocId))
	}

	if bson.IsObjectIdHex(createRegistrantForm.EmpPrefOverseasLocId) {
		registrant.Employment.PrefOverseasLoc = model.CountryById(bson.ObjectIdHex(createRegistrantForm.EmpPrefOverseasLocId))
	}

	if bson.IsObjectIdHex(createRegistrantForm.BasicInfoReligionId) {
		registrant.BasicInfo.Religion = model.ReligionById(bson.ObjectIdHex(createRegistrantForm.BasicInfoReligionId))
	}

	if bson.IsObjectIdHex(createRegistrantForm.BasicInfoCityMunId) {
		registrant.BasicInfo.CityMun = model.CityMunById(bson.ObjectIdHex(createRegistrantForm.BasicInfoCityMunId))
	}

	if createRegistrantForm.BasicInfoAge != 0 {
		registrant.BasicInfo.Age = createRegistrantForm.BasicInfoAge
	}

	if createRegistrantForm.BasicInfoWeight != 0 {
		registrant.BasicInfo.Weight = createRegistrantForm.BasicInfoWeight
	}

	if createRegistrantForm.BasicInfoHeightInInches != 0 {
		registrant.BasicInfo.HeightInInches = createRegistrantForm.BasicInfoHeightInInches
	}

	if createRegistrantForm.BasicInfoHeightInFeet != 0 {
		registrant.BasicInfo.HeightInFeet = createRegistrantForm.BasicInfoHeightInFeet
	}

	if bson.IsObjectIdHex(createRegistrantForm.EmpUnEmpStatId) {
		registrant.Employment.UnEmpStat = model.UnEmpStatById(bson.ObjectIdHex(createRegistrantForm.EmpUnEmpStatId))
	}

	if bson.IsObjectIdHex(createRegistrantForm.EmpTeminatedOverseasCountryId) {
		registrant.Employment.TeminatedOverseasCountry = model.CountryById(bson.ObjectIdHex(createRegistrantForm.EmpTeminatedOverseasCountryId))
	}

	if createRegistrantForm.EmpPassportNumberExpiryDate != "" {
		registrant.Employment.PassportNumberExpiryDate = helper.YearMonth(createRegistrantForm.EmpPassportNumberExpiryDate)
	}

	if createRegistrantForm.DisabIsDisabled {
		registrant.Disab = &model.Disab{
			IsDisabled: createRegistrantForm.DisabIsDisabled,
			Name:       model.DisabilityById(bson.ObjectIdHex(createRegistrantForm.DisabId)),
			Other:      strings.ToUpper(createRegistrantForm.DisabOther),
		}
	}

	if len(createRegistrantForm.LangIds) != 0 {
		for _, langId := range createRegistrantForm.LangIds {
			registrant.Langs = append(registrant.Langs, model.LanguageById(bson.ObjectIdHex(langId)))
		}
	}

	if len(createRegistrantForm.EmpPrefOccIds) != 0 {
		for _, empPrefOccId := range createRegistrantForm.EmpPrefOccIds {
			registrant.Employment.PrefOccs = append(registrant.Employment.PrefOccs, model.PositionById(bson.ObjectIdHex(empPrefOccId)))
		}
	}

	if len(createRegistrantForm.OtherSkillIds) != 0 {
		for _, otherSkillId := range createRegistrantForm.OtherSkillIds {
			registrant.OtherSkills = append(registrant.OtherSkills, model.OtherSkillById(bson.ObjectIdHex(otherSkillId)))
		}
	}
	formalEduArr := []model.FormalEduArr{}

	if err := json.Unmarshal([]byte(createRegistrantForm.FormalEduJson), &formalEduArr); err != nil {
		panic(err)
	}

	if len(formalEduArr) != 0 {
		for _, formalEduObj := range formalEduArr {
			formalEdu := &model.FormalEdu{
				HighestGradeCompleted: model.EduLevelById(bson.ObjectIdHex(formalEduObj.HighestGradeCompletedId)),
				CourseDegree:          model.CourseById(bson.ObjectIdHex(formalEduObj.CourseDegreeId)),
				SchoolUnivOther:       strings.ToUpper(formalEduObj.SchoolUnivOther),
				YearGrad:              helper.Year(strconv.Itoa(formalEduObj.YearGrad)),
				LastAttended:          helper.YearMonth(formalEduObj.LastAttended),
			}

			if bson.IsObjectIdHex(formalEduObj.SchoolUnivId) {
				formalEdu.SchoolUniv = model.SchoolById(bson.ObjectIdHex(formalEduObj.SchoolUnivId))
			}
			registrant.FormalEdus = append(registrant.FormalEdus, formalEdu)
		}
	}
	proLicenseArr := []model.ProLicenseArr{}

	if err := json.Unmarshal([]byte(createRegistrantForm.ProLicenseJson), &proLicenseArr); err != nil {
		panic(err)
	}

	if len(proLicenseArr) != 0 {
		for _, proLicenseObj := range proLicenseArr {
			registrant.ProLicenses = append(registrant.ProLicenses, &model.ProLicense{
				Title:      model.LicenseById(bson.ObjectIdHex(proLicenseObj.TitleId)),
				ExpiryDate: helper.YearMonth(proLicenseObj.ExpiryDate),
			})
		}
	}
	eligArr := []model.EligArr{}

	if err := json.Unmarshal([]byte(createRegistrantForm.EligJson), &eligArr); err != nil {
		panic(err)
	}

	if len(eligArr) != 0 {
		for _, eligObj := range eligArr {
			registrant.Eligs = append(registrant.Eligs, &model.Elig{
				Title:     model.EligibilityById(bson.ObjectIdHex(eligObj.TitleId)),
				YearTaken: helper.YearMonth(eligObj.YearTaken),
			})
		}
	}
	trainingArr := []model.TrainingArr{}

	if err := json.Unmarshal([]byte(createRegistrantForm.TrainingJson), &trainingArr); err != nil {
		panic(err)
	}

	if len(trainingArr) != 0 {
		for _, trainingObj := range trainingArr {
			training := &model.Training{
				Name:                strings.ToUpper(trainingObj.Name),
				SkillsAcquired:      strings.ToUpper(trainingObj.SkillsAcquired),
				PeriodOfTrainingExp: strings.ToUpper(trainingObj.PeriodOfTrainingExp),
				CertReceived:        strings.ToUpper(trainingObj.CertReceived),
				IssuingSchoolAgency: strings.ToUpper(trainingObj.IssuingSchoolAgency),
			}
			registrant.Trainings = append(registrant.Trainings, training)
		}
	}
	certArr := []model.CertArr{}

	if err := json.Unmarshal([]byte(createRegistrantForm.CertJson), &certArr); err != nil {
		panic(err)
	}

	if len(certArr) != 0 {
		for _, certObj := range certArr {
			cert := &model.Cert{
				Title:      model.CertificateById(bson.ObjectIdHex(certObj.TitleId)),
				Rating:     strings.ToUpper(certObj.Rating),
				IssuedBy:   strings.ToUpper(certObj.IssuedBy),
				DateIssued: helper.YearMonth(certObj.DateIssued),
			}
			registrant.Certs = append(registrant.Certs, cert)
		}
	}
	workExpArr := []model.WorkExpArr{}

	if err := json.Unmarshal([]byte(createRegistrantForm.WorkExpJson), &workExpArr); err != nil {
		panic(err)
	}

	if len(workExpArr) != 0 {
		for _, workExpObj := range workExpArr {
			registrant.WorkExps = append(registrant.WorkExps, &model.WorkExp{
				NameOfCompanyFirm:    strings.ToUpper(workExpObj.NameOfCompanyFirm),
				Address:              strings.ToUpper(workExpObj.Address),
				PositionHeld:         model.PositionById(bson.ObjectIdHex(workExpObj.PositionHeldId)),
				From:                 helper.YearMonth(workExpObj.From),
				To:                   helper.YearMonth(workExpObj.To),
				IsRelatedToFormalEdu: workExpObj.IsRelatedToFormalEdu,
			})
		}
	}
	registrant.Create()
	http.Redirect(w, r, "/registrants", http.StatusFound)
}

func RegistrantEmailTaken(w http.ResponseWriter, r *http.Request) {
	if taken := model.RegistrantEmailTaken(r.URL.Query().Get("basicInfoEmail")); taken {
		data := make(map[string]string)
		data["error"] = lang.Get("emailTaken")
		rd.JSON(w, http.StatusNotFound, data)
		return
	}
	w.WriteHeader(http.StatusOK)
}
