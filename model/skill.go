package model

import (
	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Skill struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Value string        `json:"value" bson:"value"`
}

func (skill *Skill) Create() *Skill {
	if err := db.C("skills").Insert(skill); err != nil {
		panic(err)
	}
	defer db.Close()
	return skill
}

func (skill Skill) Index(q string) []Skill {
	skills := []Skill{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"name": regex}

	if err := db.C("skills").Find(query).All(&skills); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return skills
}
