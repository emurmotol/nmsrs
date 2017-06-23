package model

import (
	"encoding/json"
	"io/ioutil"

	"github.com/emurmotol/nmsrs/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func otherSkillSeeder() {
	data, err := ioutil.ReadFile("import/no$oid/otherSkills.json")

	if err != nil {
		panic(err)
	}
	otherSkills := []OtherSkill{}

	if err := json.Unmarshal(data, &otherSkills); err != nil {
		panic(err)
	}
	// todo: insert to db
}

type OtherSkill struct {
	Id   bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
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
	query := bson.M{"name": regex}

	if err := db.C("otherSkills").Find(query).All(&otherSkills); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return otherSkills
}
