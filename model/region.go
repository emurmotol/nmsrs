package model

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/emurmotol/nmsrs/db"

	"gopkg.in/mgo.v2/bson"
)

type Region struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Code     string        `json:"code" bson:"code"`
	Desc     string        `json:"desc" bson:"desc"`
	PsgcCode string        `json:"psgc_code" bson:"psgcCode"`
}

type RefRegion struct {
	PsgcCode string `json:"psgcCode"`
	RegDesc  string `json:"regDesc"`
	RegCode  string `json:"regCode"`
}

func regionSeeder() {
	data, err := ioutil.ReadFile("model/data/refregion.json")

	if err != nil {
		panic(err)
	}
	refRegions := []RefRegion{}

	if err := json.Unmarshal(data, &refRegions); err != nil {
		panic(err)
	}

	for _, refRegion := range refRegions {
		region := Region{
			Id:       bson.NewObjectId(),
			Code:     refRegion.RegCode,
			Desc:     strings.ToUpper(refRegion.RegDesc),
			PsgcCode: refRegion.PsgcCode,
		}
		region.Create()
	}
}

func (region *Region) Create() *Region {
	if err := db.C("regions").Insert(region); err != nil {
		panic(err)
	}
	defer db.Close()
	return region
}
