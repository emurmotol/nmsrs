package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type EmpStat struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}

func empStatSeeder() {
	data := []string{
		"WAGED EMPLOYED",
		"SELF EMPLOYED",
		"UNEMPLOYED",
	}

	for _, name := range data {
		empStat := EmpStat{
			Id:   bson.NewObjectId(),
			Name: strings.ToUpper(name),
		}
		empStat.Create()
	}
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

	if err := db.C("empStats").Find(nil).All(&empStats); err != mgo.ErrNotFound {
		panic(err)
	} else if err == mgo.ErrNotFound {
		return nil
	}
	defer db.Close()
	return empStats
}
