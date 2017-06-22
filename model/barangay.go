package model

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"log"

	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/helper"
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
	var barangays []interface{}

	for index, refBarangay := range refBarangays {
		barangays = append(barangays, Barangay{
			Id:          bson.NewObjectId(),
			Code:        refBarangay.BrgyCode,
			Desc:        strings.ToUpper(refBarangay.BrgyDesc),
			RegCode:     refBarangay.RegCode,
			ProvCode:    refBarangay.ProvCode,
			CityMunCode: refBarangay.CityMunCode,
		})
		log.Println("refBarangay", index)
	}

	for index, chunk := range helper.ChunkSlice(barangays, 500) {
		db.C("barangays").Insert(chunk)
		log.Println("chunk", index)
	}
}

func (barangay *Barangay) Create() *Barangay {
	db.C("barangays").Insert(barangay)
	defer db.Close()
	return barangay
}
