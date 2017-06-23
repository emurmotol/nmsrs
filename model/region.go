package model

import (
	"encoding/json"
	"io/ioutil"

	"github.com/emurmotol/nmsrs/db"

	"gopkg.in/mgo.v2/bson"
)

type Region struct {
	Id       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Code     string        `json:"code" bson:"code"`
	Desc     string        `json:"desc" bson:"desc"`
	PsgcCode string        `json:"psgcCode" bson:"psgcCode"`
}

func regionSeeder() {
	data, err := ioutil.ReadFile("import/no$oid/regions.json")

	if err != nil {
		panic(err)
	}
	regions := []Region{}

	if err := json.Unmarshal(data, &regions); err != nil {
		panic(err)
	}
	// todo: insert to db
}

func (region *Region) Create() *Region {
	if err := db.C("regions").Insert(region); err != nil {
		panic(err)
	}
	defer db.Close()
	return region
}
