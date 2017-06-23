package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
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
	if err := db.C("unEmpStats").Insert(unEmpStat); err != nil {
		panic(err)
	}
	defer db.Close()
	return unEmpStat
}

func UnEmpStats() []UnEmpStat {
	unEmpStats := []UnEmpStat{}

	if err := db.C("sexes").Find(nil).All(&unEmpStats); err != mgo.ErrNotFound {
		panic(err)
	} else if err == mgo.ErrNotFound {
		return nil
	}
	defer db.Close()
	return unEmpStats
}
