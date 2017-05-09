package license

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type License struct {
	ObjectID   bson.ObjectId `schema:"_id" json:"_id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]License, error) {
	licns := []License{}

	if err := db.Licenses.Find(nil).Sort("+name").All(&licns); err != nil {
		return nil, err
	}
	return licns, nil
}

func (licn *License) Insert() (string, error) {
	licn.ObjectID = bson.NewObjectId()

	if err := db.Licenses.Insert(licn); err != nil {
		return "", err
	}
	return licn.ObjectID.Hex(), nil
}

func FindByID(id string) (*License, error) {
	var licn License

	if !bson.IsObjectIdHex(id) {
		return &licn, models.ErrInvalidObjectID
	}

	if err := db.Licenses.FindId(bson.ObjectIdHex(id)).One(&licn); err != nil {
		return &licn, err
	}
	return &licn, nil
}
