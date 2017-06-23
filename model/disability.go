package model

import (
	"encoding/json"
	"io/ioutil"

	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Disability struct {
	Id   bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}

func disabilitySeeder() {
	data, err := ioutil.ReadFile("import/no$oid/disabilities.json")

	if err != nil {
		panic(err)
	}
	disabilities := []Disability{}

	if err := json.Unmarshal(data, &disabilities); err != nil {
		panic(err)
	}
	// todo: insert to db
}

func (disability *Disability) Create() *Disability {
	if err := db.C("disabilities").Insert(disability); err != nil {
		panic(err)
	}
	defer db.Close()
	return disability
}

func Disabilities() []Disability {
	disabilities := []Disability{}

	if err := db.C("disabilities").Find(nil).All(&disabilities); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return disabilities
}
