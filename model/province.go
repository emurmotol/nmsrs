package model

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/emurmotol/nmsrs/db"

	"gopkg.in/mgo.v2/bson"
)

type Province struct {
	Id       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Code     string        `json:"code" bson:"code"`
	Desc     string        `json:"desc" bson:"desc"`
	PsgcCode string        `json:"psgcCode" bson:"psgcCode"`
	RegCode  string        `json:"regCode" bson:"regCode"`
}

func provinceSeeder() {
	data, err := ioutil.ReadFile("import/provinces.json")

	if err != nil {
		panic(err)
	}
	provinces := []Province{}

	if err := json.Unmarshal(data, &provinces); err != nil {
		panic(err)
	}
	log.Println("provinceSeeder: todo")
}

func (province *Province) Create() *Province {
	if err := db.C("provinces").Insert(province); err != nil {
		panic(err)
	}
	defer db.Close()
	return province
}
