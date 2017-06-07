package model

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type Province struct {
	ID       uint     `json:"id"`
	Code     string   `json:"code"`
	Desc     string   `json:"desc"`
	PsgcCode string   `json:"psgc_code"`
	Regions  []Region `gorm:"ForeignKey:Code;AssociationForeignKey:RegCode"`
	RegCode  string   `json:"reg_code"`
}

type RefProvince struct {
	PsgcCode string `json:"psgcCode"`
	ProvDesc string `json:"provDesc"`
	RegCode  string `json:"regCode"`
	ProvCode string `json:"provCode"`
}

func provinceSeeder() {
	data, err := ioutil.ReadFile("model/data/refprovince.json")

	if err != nil {
		panic(err)
	}
	refProvinces := []RefProvince{}

	if err := json.Unmarshal(data, &refProvinces); err != nil {
		panic(err)
	}

	for _, refProvince := range refProvinces {
		province := Province{
			Code:     refProvince.ProvCode,
			Desc:     strings.ToUpper(refProvince.ProvDesc),
			PsgcCode: refProvince.PsgcCode,
			RegCode:  refProvince.RegCode,
		}
		province.Create()
	}
}

func (province *Province) Create() *Province {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&province).Error; err != nil {
		panic(err)
	}
	return province
}
