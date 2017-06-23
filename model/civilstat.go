package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
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
	if err := db.C("civilStats").Insert(civilStat); err != nil {
		panic(err)
	}
	defer db.Close()
	return civilStat
}

func CivilStats() []CivilStat {
	civilStats := []CivilStat{}

	if err := db.C("civilStats").Find(nil).All(&civilStats); err != mgo.ErrNotFound {
		panic(err)
	} else if err == mgo.ErrNotFound {
		return nil
	}
	defer db.Close()
	return civilStats
}
