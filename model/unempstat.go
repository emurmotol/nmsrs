package model

import (
	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UnEmpStat struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
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

func UnEmpStatById(id bson.ObjectId) *UnEmpStat {
	unEmpStat := new(UnEmpStat)

	if err := db.C("unEmpStats").FindId(id).One(&unEmpStat); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return unEmpStat
}
