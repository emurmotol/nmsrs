package model

import (
	"strings"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/emurmotol/nmsrs/db"
)

func eduLevelSeeder() {
	data := []string{
		"1ST YEAR COLLEGE LEVEL",
		"1ST YEAR HIGH SCHOOL/GRADE VII (FOR K TO 12)",
		"2ND YEAR COLLEGE LEVEL",
		"2ND YEAR HIGH SCHOOL/GRADE VIII (FOR K TO 12)",
		"3RD YEAR COLLEGE LEVEL",
		"3RD YEAR HIGH SCHOOL/GRADE IX (FOR K TO 12)",
		"4TH YEAR COLLEGE LEVEL",
		"4TH YEAR HIGH SCHOOL/GRADE X (FOR K TO 12)",
		"5TH YEAR COLLEGE LEVEL",
		"COLLEGE GRADUATE",
		"ELEMENTARY GRADUATE",
		"GRADE I",
		"GRADE II",
		"GRADE III",
		"GRADE IV",
		"GRADE V",
		"GRADE VI",
		"GRADE VII",
		"GRADE VIII",
		"GRADE XI (FOR K TO 12)",
		"GRADE XII (FOR K TO 12)",
		"HIGH SCHOOL GRADUATE",
		"MASTERAL/POST GRADUATE LEVEL",
		"MASTERAL/POST GRADUATE",
		"VOCATIONAL GRADUATE",
		"VOCATIONAL UNDERGRADUATE",
	}

	for _, name := range data {
		eduLevel := EduLevel{
			Id:   bson.NewObjectId(),
			Name: strings.ToUpper(name),
		}
		eduLevel.Create()
	}
}

type EduLevel struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}

func (eduLevel *EduLevel) Create() *EduLevel {
	if err := db.C("eduLevels").Insert(eduLevel); err != nil {
		panic(err)
	}
	defer db.Close()
	return eduLevel
}

func (eduLevel EduLevel) Index(q string) []EduLevel {
	eduLevels := []EduLevel{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"name": regex}

	if err := db.C("eduLevels").Find(query).All(&eduLevels); err != mgo.ErrNotFound {
		panic(err)
	} else if err == mgo.ErrNotFound {
		return nil
	}
	defer db.Close()
	return eduLevels
}
