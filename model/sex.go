package model

import (
	"encoding/json"
	"io/ioutil"

	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Sex struct {
	Id   bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}

func sexSeeder() {
	data, err := ioutil.ReadFile("import/no$oid/sexes.json")

	if err != nil {
		panic(err)
	}
	sexes := []Sex{}

	if err := json.Unmarshal(data, &sexes); err != nil {
		panic(err)
	}
	// todo: insert to db
}

func (sex *Sex) Create() *Sex {
	if err := db.C("sexes").Insert(sex); err != nil {
		panic(err)
	}
	defer db.Close()
	return sex
}

func Sexes() []Sex {
	sexes := []Sex{}

	if err := db.C("sexes").Find(nil).All(&sexes); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return sexes
}
