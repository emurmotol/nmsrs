package model

import (
	"encoding/json"
	"io/ioutil"

	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func countrySeeder() {
	data, err := ioutil.ReadFile("import/no$oid/countries.json")

	if err != nil {
		panic(err)
	}
	countries := []Country{}

	if err := json.Unmarshal(data, &countries); err != nil {
		panic(err)
	}
	// todo: insert to db
}

type Country struct {
	Id   bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}

func (country *Country) Create() *Country {
	if err := db.C("countries").Insert(country); err != nil {
		panic(err)
	}
	defer db.Close()
	return country
}

func (country Country) Index(q string) []Country {
	countries := []Country{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"name": regex}

	if err := db.C("countries").Find(query).All(&countries); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return countries
}
