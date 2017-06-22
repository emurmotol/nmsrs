package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/db"

	"gopkg.in/mgo.v2/bson"
)

type CivilStat struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}

func civilStatSeeder() {
	data := []string{
		"SINGLE",
		"WIdOWED",
		"MARRIED",
		"SEPARATED",
		"OTHER",
	}

	for _, name := range data {
		civilStat := CivilStat{
			Id:   bson.NewObjectId(),
			Name: strings.ToUpper(name),
		}
		civilStat.Create()
	}
}

func (civilStat *CivilStat) Create() *CivilStat {
	db.C("civilStats").Insert(civilStat)
	defer db.Close()
	return civilStat
}

func CivilStats() []CivilStat {
	civilStats := []CivilStat{}
	db.C("civilStats").Find(nil).All(&civilStats)
	defer db.Close()
	return civilStats
}
