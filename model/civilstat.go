package model

import (
	"encoding/json"
	"io/ioutil"

	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CivilStat struct {
	Id   bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}

func civilStatSeeder() {
	data, err := ioutil.ReadFile("import/no$oid/civilStats.json")

	if err != nil {
		panic(err)
	}
	civilStats := []CivilStat{}

	if err := json.Unmarshal(data, &civilStats); err != nil {
		panic(err)
	}
	// todo: insert to db
}

func (civilStat *CivilStat) Create() *CivilStat {
	if err := db.C("civilStats").Insert(civilStat); err != nil {
		panic(err)
	}
	defer db.Close()
	return civilStat
}

func CivilStats() []CivilStat {
	civilStats := []CivilStat{}

	if err := db.C("civilStats").Find(nil).All(&civilStats); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return civilStats
}
