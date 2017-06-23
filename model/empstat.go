package model

import (
	"encoding/json"
	"io/ioutil"

	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type EmpStat struct {
	Id   bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}

func empStatSeeder() {
	data, err := ioutil.ReadFile("import/empStats.json")

	if err != nil {
		panic(err)
	}
	empStats := []EmpStat{}

	if err := json.Unmarshal(data, &empStats); err != nil {
		panic(err)
	}
	// todo: insert to db
}

func (empStat *EmpStat) Create() *EmpStat {
	if err := db.C("empStats").Insert(empStat); err != nil {
		panic(err)
	}
	defer db.Close()
	return empStat
}

func EmpStats() []EmpStat {
	empStats := []EmpStat{}

	if err := db.C("empStats").Find(nil).All(&empStats); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return empStats
}
