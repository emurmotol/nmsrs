package model

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	
)

type EduLevel struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Value string        `json:"value" bson:"value"`
}

func (eduLevel *EduLevel) Create() *EduLevel {
	if err := db.C("eduLevels").Insert(eduLevel); err != nil {
		panic(err)
	}
	return eduLevel
}

func (eduLevel EduLevel) Index(q string) []EduLevel {
	eduLevels := []EduLevel{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"value": regex}

	if err := db.C("eduLevels").Find(query).All(&eduLevels); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return eduLevels
}

func EduLevelById(id bson.ObjectId) *EduLevel {
	eduLevel := new(EduLevel)

	if err := db.C("eduLevels").FindId(id).One(&eduLevel); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return eduLevel
}
