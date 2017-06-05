package model

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/emurmotol/nmsrs/database"
)

type Province struct {
	ID       int    `json:"id"`
	Code     string `json:"code"`
	Desc     string `json:"desc"`
	PsgcCode string `json:"psgc_code"`
	RegCode  string `json:"reg_code"`
}

type RefProvince struct {
	PsgcCode string `json:"psgcCode"`
	ProvDesc string `json:"provDesc"`
	RegCode  string `json:"regCode"`
	ProvCode string `json:"provCode"`
}

func ProvinceSeeder() {
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

		if _, err := province.Create(); err != nil {
			panic(err)
		}
	}
}

func (province *Province) Create() (*Province, error) {
	db := database.Conn()
	defer db.Close()

	if err := db.Create(&province).Error; err != nil {
		return nil, err
	}
	return province, nil
}

func (province Province) Search(q string) []Province {
	db := database.Conn()
	defer db.Close()

	provinces := []Province{}
	results := make(chan []Province)

	go func() {
		db.Find(&provinces, "name LIKE ?", database.WrapLike(q))
		results <- provinces
	}()
	return <-results
}
