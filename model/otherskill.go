package model

import (
	"github.com/emurmotol/nmsrs/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type OtherSkill struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Value string        `json:"value" bson:"value"`
}

func (otherSkill *OtherSkill) Create() *OtherSkill {
	if err := db.C("otherSkills").Insert(otherSkill); err != nil {
		panic(err)
	}
	defer db.Close()
	return otherSkill
}

func (otherSkill OtherSkill) Index(q string) []OtherSkill {
	otherSkills := []OtherSkill{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"value": regex}

	if err := db.C("otherSkills").Find(query).All(&otherSkills); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return otherSkills
}

func OtherSkillById(id bson.ObjectId) *OtherSkill {
	otherSkill := new(OtherSkill)

	if err := db.C("otherSkills").FindId(id).One(&otherSkill); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return otherSkill
}
