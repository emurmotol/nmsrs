package model

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type Barangay struct {
	Code        string `gorm:"primary_key" json:"code"`
	Desc        string `json:"desc"`
	RegCode     string `json:"reg_code"`
	ProvCode    string `json:"prov_code"`
	CityMunCode string `json:"city_mun_code"`
}

type RefBarangay struct {
	BrgyCode    string `json:"brgyCode"`
	BrgyDesc    string `json:"brgyDesc"`
	RegCode     string `json:"regCode"`
	ProvCode    string `json:"provCode"`
	CityMunCode string `json:"cityMunCode"`
}

func barangaySeeder() {
	data, err := ioutil.ReadFile("model/data/refbarangay.json")

	if err != nil {
		panic(err)
	}
	refBarangays := []RefBarangay{}

	if err := json.Unmarshal(data, &refBarangays); err != nil {
		panic(err)
	}

	for _, refBarangay := range refBarangays {
		barangay := Barangay{
			Code:        refBarangay.BrgyCode,
			Desc:        strings.ToUpper(refBarangay.BrgyDesc),
			RegCode:     refBarangay.RegCode,
			ProvCode:    refBarangay.ProvCode,
			CityMunCode: refBarangay.CityMunCode,
		}
		barangay.Create()
	}
}

func (barangay *Barangay) Create() *Barangay {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&barangay).Error; err != nil {
		panic(err)
	}
	return barangay
}
