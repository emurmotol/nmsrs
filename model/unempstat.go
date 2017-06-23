package model

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UnEmpStat struct {
	Id   bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}

func unEmpStatSeeder() {
	data, err := ioutil.ReadFile("import/no$oid/unEmpStats.json")

	if err != nil {
		panic(err)
	}
	unEmpStats := []UnEmpStat{}

	if err := json.Unmarshal(data, &unEmpStats); err != nil {
		panic(err)
	}
	log.Println("unEmpStatSeeder: todo")
}

func (unEmpStat *UnEmpStat) Create() *UnEmpStat {
	if err := db.C("unEmpStats").Insert(unEmpStat); err != nil {
		panic(err)
	}
	defer db.Close()
	return unEmpStat
}

func UnEmpStats() []UnEmpStat {
	unEmpStats := []UnEmpStat{}

	if err := db.C("unEmpStats").Find(nil).All(&unEmpStats); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return unEmpStats
}
