package controller

import (
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"github.com/emurmotol/nmsrs/model"
	"github.com/pressly/chi"
)

func CertificateIndex(w http.ResponseWriter, r *http.Request) {
	certificate := model.Certificate{}
	rd.JSON(w, http.StatusOK, certificate.Index(r.URL.Query().Get("q")))
}

func CityMunBarangayIndex(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "cityMunId"))

	if err != nil {
		panic(err)
	}
	cityMun := model.CityMunById(bson.ObjectId(id))
	rd.JSON(w, http.StatusOK, cityMun.BarangayIndex(r.URL.Query().Get("q")))
}

func CityMunProvinceIndex(w http.ResponseWriter, r *http.Request) {
	cityMun := model.CityMun{}
	rd.JSON(w, http.StatusOK, cityMun.ProvinceIndex(r.URL.Query().Get("q")))
}

func CountryIndex(w http.ResponseWriter, r *http.Request) {
	country := model.Country{}
	rd.JSON(w, http.StatusOK, country.Index(r.URL.Query().Get("q")))
}

func CourseIndex(w http.ResponseWriter, r *http.Request) {
	course := model.Course{}
	rd.JSON(w, http.StatusOK, course.Index(r.URL.Query().Get("q")))
}

func EduLevelIndex(w http.ResponseWriter, r *http.Request) {
	eduLevel := model.EduLevel{}
	rd.JSON(w, http.StatusOK, eduLevel.Index(r.URL.Query().Get("q")))
}

func EligibilityIndex(w http.ResponseWriter, r *http.Request) {
	eligibility := model.Eligibility{}
	rd.JSON(w, http.StatusOK, eligibility.Index(r.URL.Query().Get("q")))
}

func IndustryIndex(w http.ResponseWriter, r *http.Request) {
	industry := model.Industry{}
	rd.JSON(w, http.StatusOK, industry.Index(r.URL.Query().Get("q")))
}

func LanguageIndex(w http.ResponseWriter, r *http.Request) {
	language := model.Language{}
	rd.JSON(w, http.StatusOK, language.Index(r.URL.Query().Get("q")))
}

func LicenseIndex(w http.ResponseWriter, r *http.Request) {
	license := model.License{}
	rd.JSON(w, http.StatusOK, license.Index(r.URL.Query().Get("q")))
}

func OtherSkillIndex(w http.ResponseWriter, r *http.Request) {
	otherSkill := model.OtherSkill{}
	rd.JSON(w, http.StatusOK, otherSkill.Index(r.URL.Query().Get("q")))
}

func PositionIndex(w http.ResponseWriter, r *http.Request) {
	position := model.Position{}
	rd.JSON(w, http.StatusOK, position.Index(r.URL.Query().Get("q")))
}

func ReligionIndex(w http.ResponseWriter, r *http.Request) {
	religion := model.Religion{}
	rd.JSON(w, http.StatusOK, religion.Index(r.URL.Query().Get("q")))
}

func SchoolIndex(w http.ResponseWriter, r *http.Request) {
	school := model.School{}
	rd.JSON(w, http.StatusOK, school.Index(r.URL.Query().Get("q")))
}

func SkillIndex(w http.ResponseWriter, r *http.Request) {
	skill := model.Skill{}
	rd.JSON(w, http.StatusOK, skill.Index(r.URL.Query().Get("q")))
}

func UnEmpStatIndex(w http.ResponseWriter, r *http.Request) {
	rd.JSON(w, http.StatusOK, model.UnEmpStats())
}
