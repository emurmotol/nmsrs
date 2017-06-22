package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/db"

	"gopkg.in/mgo.v2/bson"
)

type UnEmpStat struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}

func unEmpStatSeeder() {
	data := []string{
		"NEW ENTRANT/FRESH GRADUATE",
		"FINISHED CONTRACT",
		"RESIGNED",
		"TERMINATED/LAID OFF, LOCAL",
		"TERMINATED/LAID OFF, OVERSEAS",
	}

	for _, name := range data {
		unEmpStat := UnEmpStat{
			Id:   bson.NewObjectId(),
			Name: strings.ToUpper(name),
		}
		unEmpStat.Create()
	}
}

func (unEmpStat *UnEmpStat) Create() *UnEmpStat {
	db.C("unEmpStats").Insert(unEmpStat)
	defer db.Close()
	return unEmpStat
}

func UnEmpStats() []UnEmpStat {
	unEmpStats := []UnEmpStat{}
	db.C("sexes").Find(nil).All(&unEmpStats)
	defer db.Close()
	return unEmpStats
}
