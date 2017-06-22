package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/db"

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
	db.C("empStats").Insert(empStat)
	defer db.Close()
	return empStat
}

func EmpStats() []EmpStat {
	empStats := []EmpStat{}
	db.C("empStats").Find(nil).All(&empStats)
	defer db.Close()
	return empStats
}
