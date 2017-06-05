package model

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type Barangay struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
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

func BarangaySeeder() {
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

		if _, err := barangay.Create(); err != nil {
			panic(err)
		}
	}
}

func (barangay *Barangay) Create() (*Barangay, error) {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&barangay).Error; err != nil {
		return nil, err
	}
	return barangay, nil
}
