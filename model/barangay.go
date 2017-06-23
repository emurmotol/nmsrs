package model

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type Barangay struct {
	Id          bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Code        string        `json:"code" bson:"code"`
	Desc        string        `json:"desc" bson:"desc"`
	RegCode     string        `json:"regCode" bson:"regCode"`
	ProvCode    string        `json:"provCode" bson:"provCode"`
	CityMunCode string        `json:"cityMunCode" bson:"cityMunCode"`
}

func barangaySeeder() {
	data, err := ioutil.ReadFile("import/no$oid/barangays.json")

	if err != nil {
		panic(err)
	}
	barangays := []Barangay{}

	if err := json.Unmarshal(data, &barangays); err != nil {
		panic(err)
	}
	log.Println("barangaySeeder: todo")
}

func (barangay *Barangay) Create() *Barangay {
	if err := db.C("barangays").Insert(barangay); err != nil {
		panic(err)
	}
	defer db.Close()
	return barangay
}
