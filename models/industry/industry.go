package industry

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Industry struct {
	ID   bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Industry, error) {
	inds := []Industry{}

	if err := db.Industries.Find(bson.M{}).Sort("+name").All(&inds); err != nil {
		return nil, err
	}
	return inds, nil
}

func (ind *Industry) Insert() (string, error) {
	ind.ID = bson.NewObjectId()

	if err := db.Industries.Insert(ind); err != nil {
		return "", err
	}
	return ind.ID.Hex(), nil
}

func Find(id string) (*Industry, error) {
	var ind Industry

	if !bson.IsObjectIdHex(id) {
		return &ind, models.ErrInvalidObjectID
	}

	if err := db.Industries.FindId(bson.ObjectIdHex(id)).One(&ind); err != nil {
		return &ind, err
	}
	return &ind, nil
}
