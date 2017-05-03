package course

import (
	"github.com/zneyrl/nmsrs/db"
	"github.com/zneyrl/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Course struct {
	ID   bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Course, error) {
	cours := []Course{}

	if err := db.Courses.Find(bson.M{}).Sort("+name").All(&cours); err != nil {
		return nil, err
	}
	return cours, nil
}

func (cour *Course) Insert() (string, error) {
	cour.ID = bson.NewObjectId()

	if err := db.Courses.Insert(cour); err != nil {
		return "", err
	}
	return cour.ID.Hex(), nil
}

func Find(id string) (*Course, error) {
	var cour Course

	if !bson.IsObjectIdHex(id) {
		return &cour, models.ErrInvalidObjectID
	}

	if err := db.Courses.FindId(bson.ObjectIdHex(id)).One(&cour); err != nil {
		return &cour, err
	}
	return &cour, nil
}
