package model

import (
	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type EmpStat struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
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

func EmpStatById(id bson.ObjectId) *EmpStat {
	empStat := new(EmpStat)

	if err := db.C("empStats").FindId(id).One(&empStat); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return empStat
}
