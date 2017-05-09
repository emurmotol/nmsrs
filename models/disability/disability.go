package disability

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Disability struct {
	ObjectID   bson.ObjectId `schema:"_id" json:"_id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Disability, error) {
	disabs := []Disability{}

	if err := db.Disabilities.Find(nil).All(&disabs); err != nil {
		return nil, err
	}
	return disabs, nil
}

func (disab *Disability) Insert() (string, error) {
	disab.ObjectID = bson.NewObjectId()

	if err := db.Disabilities.Insert(disab); err != nil {
		return "", err
	}
	return disab.ObjectID.Hex(), nil
}

func FindByID(id string) (*Disability, error) {
	var disab Disability

	if !bson.IsObjectIdHex(id) {
		return &disab, models.ErrInvalidObjectID
	}

	if err := db.Disabilities.FindId(bson.ObjectIdHex(id)).One(&disab); err != nil {
		return &disab, err
	}
	return &disab, nil
}
