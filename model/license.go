package model

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func licenseSeeder() {
	data, err := ioutil.ReadFile("import/no$oid/licenses.json")

	if err != nil {
		panic(err)
	}
	licenses := []License{}

	if err := json.Unmarshal(data, &licenses); err != nil {
		panic(err)
	}
	log.Println("licenseSeeder: todo")
}

type License struct {
	Id   bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}

func (license *License) Create() *License {
	if err := db.C("licenses").Insert(license); err != nil {
		panic(err)
	}
	defer db.Close()
	return license
}

func (license License) Index(q string) []License {
	licenses := []License{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"name": regex}

	if err := db.C("licenses").Find(query).All(&licenses); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return licenses
}
