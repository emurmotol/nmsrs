package model

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type Barangay struct {
	Id          bson.ObjectId `json:"id" bson:"_id"`
	Code        string        `json:"code" bson:"code"`
	Desc        string        `json:"desc" bson:"desc"`
	RegCode     string        `json:"reg_code" bson:"regCode"`
	ProvCode    string        `json:"prov_code" bson:"provCode"`
	CityMunCode string        `json:"city_mun_code" bson:"cityMunCode"`
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
			Id:          bson.NewObjectId(),
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
	if err := db.C("barangays").Insert(barangay); err != nil {
		panic(err)
	}
	defer db.Close()
	return barangay
}
