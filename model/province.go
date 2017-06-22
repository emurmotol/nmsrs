package model

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/emurmotol/nmsrs/db"

	"gopkg.in/mgo.v2/bson"
)

type Province struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Code     string        `json:"code" bson:"code"`
	Desc     string        `json:"desc" bson:"desc"`
	PsgcCode string        `json:"psgc_code" bson:"psgcCode"`
	RegCode  string        `json:"reg_code" bson:"regCode"`
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
			Id:       bson.NewObjectId(),
			Code:     refProvince.ProvCode,
			Desc:     strings.ToUpper(refProvince.ProvDesc),
			PsgcCode: refProvince.PsgcCode,
			RegCode:  refProvince.RegCode,
		}
		province.Create()
	}
}

func (province *Province) Create() *Province {
	db.C("provinces").Insert(province)
	defer db.Close()
	return province
}
