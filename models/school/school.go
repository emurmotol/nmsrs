package school

import (
	"github.com/zneyrl/nmsrs/db"
	"github.com/zneyrl/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type School struct {
	ID   bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]School, error) {
	schs := []School{}

	if err := db.Schools.Find(bson.M{}).Sort("+name").All(&schs); err != nil {
		return nil, err
	}
	return schs, nil
}

func (sch *School) Insert() (string, error) {
	sch.ID = bson.NewObjectId()

	if err := db.Schools.Insert(sch); err != nil {
		return "", err
	}
	return sch.ID.Hex(), nil
}

func Find(id string) (*School, error) {
	var sch School

	if !bson.IsObjectIdHex(id) {
		return &sch, models.ErrInvalidObjectID
	}

	if err := db.Schools.FindId(bson.ObjectIdHex(id)).One(&sch); err != nil {
		return &sch, err
	}
	return &sch, nil
}
