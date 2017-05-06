package skill

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Skill struct {
	ID   bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Skill, error) {
	skills := []Skill{}

	if err := db.Skills.Find(nil).Sort("+name").All(&skills); err != nil {
		return nil, err
	}
	return skills, nil
}

func (skill *Skill) Insert() (string, error) {
	skill.ID = bson.NewObjectId()

	if err := db.Skills.Insert(skill); err != nil {
		return "", err
	}
	return skill.ID.Hex(), nil
}

func FindByID(id string) (*Skill, error) {
	var skill Skill

	if !bson.IsObjectIdHex(id) {
		return &skill, models.ErrInvalidObjectID
	}

	if err := db.Skills.FindId(bson.ObjectIdHex(id)).One(&skill); err != nil {
		return &skill, err
	}
	return &skill, nil
}
