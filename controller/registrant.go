package controller

import (
	"html/template"
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"

	"github.com/emurmotol/nmsrs/helper"
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
	prefOccs := []*model.Position{}

	for _, empPrefOccId := range createRegistrantForm.EmpPrefOccIds {
		prefOccs = append(prefOccs, model.PositionById(bson.ObjectIdHex(empPrefOccId)))
	}
	langs := []*model.Language{}

	for _, langId := range createRegistrantForm.LangIds {
		langs = append(langs, model.LanguageById(bson.ObjectIdHex(langId)))
	}
	// formalEduArr := []model.FormalEduArr{}

	// if err := json.Unmarshal([]byte(createRegistrantForm.FormalEduJson), &formalEduArr); err != nil {
	// 	panic(err)
	// }
	// formalEdus := []*model.FormalEdu{}

	// for _, formalEduObj := range formalEduArr {
	// 	formalEdus = append(formalEdus, &model.FormalEdu{
	// 		HighestGradeCompleted: model.EduLevelById(bson.ObjectIdHex(formalEduObj.HighestGradeCompletedId)),
	// 		CourseDegree:          model.CourseById(bson.ObjectIdHex(formalEduObj.CourseDegreeId)),
	// 		SchoolUniv:            model.SchoolById(bson.ObjectIdHex(formalEduObj.SchoolUnivId)),
	// 		SchoolUnivOther:       formalEduObj.SchoolUnivOther,
	// 		YearGrad:              helper.YearMonth(formalEduObj.YearGrad),
	// 		LastAttended:          helper.YearMonth(formalEduObj.LastAttended),
	// 	})
	// }
	// proLicenseArr := []model.ProLicenseArr{}

	// if err := json.Unmarshal([]byte(createRegistrantForm.ProLicenseJson), &proLicenseArr); err != nil {
	// 	panic(err)
	// }
	// proLicenses := []*model.ProLicense{}

	// for _, proLicenseObj := range proLicenseArr {
	// 	proLicenses = append(proLicenses, &model.ProLicense{
	// 		Title:      model.LicenseById(bson.ObjectIdHex(proLicenseObj.TitleId)),
	// 		ExpiryDate: helper.YearMonth(proLicenseObj.ExpiryDate),
	// 	})
	// }
	// eligArr := []model.EligArr{}

	// if err := json.Unmarshal([]byte(createRegistrantForm.EligJson), &eligArr); err != nil {
	// 	panic(err)
	// }
	// eligs := []*model.Elig{}

	// for _, eligObj := range eligArr {
	// 	eligs = append(eligs, &model.Elig{
	// 		Title:     model.EligibilityById(bson.ObjectIdHex(eligObj.TitleId)),
	// 		YearTaken: helper.YearMonth(eligObj.YearTaken),
	// 	})
	// }
	// trainingArr := []model.TrainingArr{}

	// if err := json.Unmarshal([]byte(createRegistrantForm.TrainingJson), &trainingArr); err != nil {
	// 	panic(err)
	// }
	// trainings := []*model.Training{}

	// for _, trainingObj := range trainingArr {
	// 	trainings = append(trainings, &model.Training{
	// 		Name:                trainingObj.Name,
	// 		SkillsAcquired:      trainingObj.SkillsAcquired,
	// 		PeriodOfTrainingExp: trainingObj.PeriodOfTrainingExp,
	// 		CertReceived:        trainingObj.CertReceived,
	// 		IssuingSchoolAgency: trainingObj.IssuingSchoolAgency,
	// 	})
	// }
	// certArr := []model.CertArr{}

	// if err := json.Unmarshal([]byte(createRegistrantForm.EligJson), &certArr); err != nil {
	// 	panic(err)
	// }
	// certs := []*model.Cert{}

	// for _, certObj := range certArr {
	// 	certs = append(certs, &model.Cert{
	// 		Title:      model.CertificateById(bson.ObjectIdHex(certObj.TitleId)),
	// 		Rating:     certObj.Rating,
	// 		IssuedBy:   certObj.IssuedBy,
	// 		DateIssued: helper.YearMonth(certObj.DateIssued),
	// 	})
	// }
	// workExpArr := []model.WorkExpArr{}

	// if err := json.Unmarshal([]byte(createRegistrantForm.WorkExpJson), &workExpArr); err != nil {
	// 	panic(err)
	// }
	// workExps := []*model.WorkExp{}

	// for _, workExpObj := range workExpArr {
	// 	workExps = append(workExps, &model.WorkExp{
	// 		NameOfCompanyFirm:    workExpObj.NameOfCompanyFirm,
	// 		Address:              workExpObj.Address,
	// 		PositionHeld:         model.PositionById(bson.ObjectIdHex(workExpObj.PositionHeldId)),
	// 		From:                 helper.YearMonth(workExpObj.From),
	// 		To:                   helper.YearMonth(workExpObj.To),
	// 		IsRelatedToFormalEdu: workExpObj.IsRelatedToFormalEdu,
	// 	})
	// }
	otherSkills := []*model.OtherSkill{}

	for _, otherSkillId := range createRegistrantForm.OtherSkillIds {
		otherSkills = append(otherSkills, model.OtherSkillById(bson.ObjectIdHex(otherSkillId)))
	}

	registrant := model.Registrant{
		RegisteredAt: helper.ShortDate(createRegistrantForm.RegisteredAt),
		IAccept:      createRegistrantForm.IAccept,
		PersonalInfo: &model.PersonalInfo{
			HasPhoto:   hasPhoto,
			FamilyName: createRegistrantForm.PersonalInfoFamilyName,
			GivenName:  createRegistrantForm.PersonalInfoGivenName,
			MiddleName: createRegistrantForm.PersonalInfoMiddleName,
			Birthdate:  helper.ShortDate(createRegistrantForm.PersonalInfoBirthdate),
			Password:   createRegistrantForm.PersonalInfoPassword,
		},
		BasicInfo: &model.BasicInfo{
			StSub:   createRegistrantForm.BasicInfoStSub,
			CityMun: model.CityMunById(bson.ObjectIdHex(createRegistrantForm.BasicInfoCityMunId)),
			// Province:       model.ProvinceById(bson.ObjectIdHex(createRegistrantForm.BasicInfoProvinceId)),
			// Barangay:       model.BarangayById(bson.ObjectIdHex(createRegistrantForm.BasicInfoBarangayId)),
			PlaceOfBirth:   createRegistrantForm.BasicInfoPlaceOfBirth,
			Religion:       model.ReligionById(bson.ObjectIdHex(createRegistrantForm.BasicInfoReligionId)),
			CivilStat:      model.CivilStatById(bson.ObjectIdHex(createRegistrantForm.BasicInfoCivilStatId)),
			CivilStatOther: createRegistrantForm.BasicInfoCivilStatOther,
			Sex:            model.SexById(bson.ObjectIdHex(createRegistrantForm.BasicInfoSexId)),
			Age:            createRegistrantForm.BasicInfoAge,
			Height:         createRegistrantForm.BasicInfoHeight,
			Weight:         createRegistrantForm.BasicInfoWeight,
			LandlineNumber: createRegistrantForm.BasicInfoLandlineNumber,
			MobileNumber:   createRegistrantForm.BasicInfoMobileNumber,
			Email:          strings.ToLower(createRegistrantForm.BasicInfoEmail),
		},
		Employment: &model.Employment{
			Stat: model.EmpStatById(bson.ObjectIdHex(createRegistrantForm.EmpStatId)),
			// UnEmpStat:                model.UnEmpStatById(bson.ObjectIdHex(createRegistrantForm.EmpUnEmpStatId)),
			// TeminatedOverseasCountry: model.CountryById(bson.ObjectIdHex(createRegistrantForm.EmpTeminatedOverseasCountryId)),
			IsActivelyLookingForWork: createRegistrantForm.EmpIsActivelyLookingForWork,
			PrefOccs:                 prefOccs,
			PrefLocalLoc:             model.CityMunById(bson.ObjectIdHex(createRegistrantForm.EmpPrefLocalLocId)),
			PrefOverseasLoc:          model.CountryById(bson.ObjectIdHex(createRegistrantForm.EmpPrefOverseasLocId)),
			PassportNumber:           createRegistrantForm.EmpPassportNumber,
			// PassportNumberExpiryDate: helper.YearMonth(createRegistrantForm.EmpPassportNumberExpiryDate),
		},
		Disab: &model.Disab{
			IsDisabled: createRegistrantForm.DisabIsDisabled,
			// Name:       model.DisabilityById(bson.ObjectIdHex(createRegistrantForm.DisabId)),
			Other: createRegistrantForm.DisabOther,
		},
		Langs: langs,
		// FormalEdus:      formalEdus,
		// ProLicenses:     proLicenses,
		// Eligs:           eligs,
		// Trainings:       trainings,
		// Certs:           certs,
		// WorkExps:        workExps,
		OtherSkills:     otherSkills,
		OtherSkillOther: createRegistrantForm.OtherSkillOther,
	}
	registrant.Create()
	http.Redirect(w, r, "/registrants/create", http.StatusFound)
}

// func RegistrantEmailTaken(w http.ResponseWriter, r *http.Request) {
// 	if taken := model.RegistrantEmailTaken(r.URL.Query().Get("email")); taken {
// 		data := make(map[string]string)
// 		data["error"] = lang.Get("emailTaken")
// 		rd.JSON(w, http.StatusNotFound, data)
// 		return
// 	}
// 	w.WriteHeader(http.StatusOK)
// }
