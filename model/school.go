package model

import (
	
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type School struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Value string        `json:"value" bson:"value"`
}

func (school *School) Create() *School {
	if err := db.C("schools").Insert(school); err != nil {
		panic(err)
	}
	return school
}

func (school School) Index(q string) []School {
	schools := []School{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"value": regex}

	if err := db.C("schools").Find(query).All(&schools); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return schools
}

func SchoolById(id bson.ObjectId) *School {
	school := new(School)

	if err := db.C("schools").FindId(id).One(&school); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return school
}
