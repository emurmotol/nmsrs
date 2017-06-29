package model

import (
	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CivilStat struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Value string        `json:"value" bson:"value"`
}

func (civilStat *CivilStat) Create() *CivilStat {
	if err := db.C("civilStats").Insert(civilStat); err != nil {
		panic(err)
	}
	defer db.Close()
	return civilStat
}

func CivilStats() []CivilStat {
	var civilStats, civilStatsArranged []CivilStat

	if err := db.C("civilStats").Find(nil).All(&civilStats); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	var civilStatOther CivilStat

	for _, civilStat := range civilStats {
		if civilStat.Id.Hex() == "594cb5fd472e11263c3291aa" {
			civilStatOther = civilStat
			continue
		}
		civilStatsArranged = append(civilStatsArranged, civilStat)
	}
	civilStatsArranged = append(civilStatsArranged, civilStatOther)
	return civilStatsArranged
}

func CivilStatById(id bson.ObjectId) *CivilStat {
	civilStat := new(CivilStat)

	if err := db.C("civilStats").FindId(id).One(&civilStat); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return civilStat
}
