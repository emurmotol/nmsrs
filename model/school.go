package model

import (
	"encoding/json"
	"io/ioutil"

	"github.com/emurmotol/nmsrs/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func schoolSeeder() {
	data, err := ioutil.ReadFile("import/no$oid/schools.json")

	if err != nil {
		panic(err)
	}
	schools := []School{}

	if err := json.Unmarshal(data, &schools); err != nil {
		panic(err)
	}
	// todo: insert to db
}

type School struct {
	Id   bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}

func (school *School) Create() *School {
	if err := db.C("schools").Insert(school); err != nil {
		panic(err)
	}
	defer db.Close()
	return school
}

func (school School) Index(q string) []School {
	schools := []School{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"name": regex}

	if err := db.C("schools").Find(query).All(&schools); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return schools
}
