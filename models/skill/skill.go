package skill

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type Skill struct {
	ID   int    `schema:"id" json:"id" bson:"id,omitempty"`
	Name string `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Skill, error) {
	skills := []Skill{}

	if err := db.Skills.Find(nil).Sort("+name").All(&skills); err != nil {
		return nil, err
	}
	return skills, nil
}

func (skill *Skill) Insert() (int, error) {
	if err := db.Skills.Insert(skill); err != nil {
		return 0, err
	}
	return skill.ID, nil
}

func FindByID(id int) (*Skill, error) {
	var skill Skill

	if err := db.Skills.Find(bson.M{"id": id}).One(&skill); err != nil {
		return &skill, err
	}
	return &skill, nil
}

func Search(query interface{}) ([]Skill, error) {
	skills := []Skill{}

	if err := db.Skills.Find(query).Sort("+name").All(&skills); err != nil {
		return nil, err
	}
	return skills, nil
}
