package educationlevel

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type EducationLevel struct {
	ObjectID   bson.ObjectId `schema:"_id" json:"_id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]EducationLevel, error) {
	edulvls := []EducationLevel{}

	if err := db.EducationLevels.Find(nil).Sort("+name").All(&edulvls); err != nil {
		return nil, err
	}
	return edulvls, nil
}

func (edulvl *EducationLevel) Insert() (string, error) {
	edulvl.ObjectID = bson.NewObjectId()

	if err := db.EducationLevels.Insert(edulvl); err != nil {
		return "", err
	}
	return edulvl.ObjectID.Hex(), nil
}

func FindByID(id string) (*EducationLevel, error) {
	var edulvl EducationLevel

	if !bson.IsObjectIdHex(id) {
		return &edulvl, models.ErrInvalidObjectID
	}

	if err := db.EducationLevels.FindId(bson.ObjectIdHex(id)).One(&edulvl); err != nil {
		return &edulvl, err
	}
	return &edulvl, nil
}
