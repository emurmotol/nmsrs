package course

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Course struct {
	ObjectID   bson.ObjectId `schema:"_id" json:"_id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Course, error) {
	cours := []Course{}

	if err := db.Courses.Find(nil).Sort("+name").All(&cours); err != nil {
		return nil, err
	}
	return cours, nil
}

func (cour *Course) Insert() (string, error) {
	cour.ObjectID = bson.NewObjectId()

	if err := db.Courses.Insert(cour); err != nil {
		return "", err
	}
	return cour.ObjectID.Hex(), nil
}

func FindByID(id string) (*Course, error) {
	var cour Course

	if !bson.IsObjectIdHex(id) {
		return &cour, models.ErrInvalidObjectID
	}

	if err := db.Courses.FindId(bson.ObjectIdHex(id)).One(&cour); err != nil {
		return &cour, err
	}
	return &cour, nil
}
