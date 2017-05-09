package educationlevel

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type EducationLevel struct {
	ID   int    `schema:"id" json:"id" bson:"id,omitempty"`
	Name string `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]EducationLevel, error) {
	edulvls := []EducationLevel{}

	if err := db.EducationLevels.Find(nil).Sort("+name").All(&edulvls); err != nil {
		return nil, err
	}
	return edulvls, nil
}

func (edulvl *EducationLevel) Insert() (int, error) {
	if err := db.EducationLevels.Insert(edulvl); err != nil {
		return 0, err
	}
	return edulvl.ID, nil
}

func FindByID(id int) (*EducationLevel, error) {
	var edulvl EducationLevel

	if err := db.EducationLevels.Find(bson.M{"id": id}).One(&edulvl); err != nil {
		return &edulvl, err
	}
	return &edulvl, nil
}

func Search(query interface{}) ([]EducationLevel, error) {
	edulvls := []EducationLevel{}

	if err := db.EducationLevels.Find(query).Sort("+name").All(&edulvls); err != nil {
		return nil, err
	}
	return edulvls, nil
}
