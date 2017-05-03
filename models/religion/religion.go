package religion

import (
	"github.com/zneyrl/nmsrs/db"
	"github.com/zneyrl/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Religion struct {
	ID   bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Religion, error) {
	religs := []Religion{}

	if err := db.Religions.Find(bson.M{}).All(&religs); err != nil {
		return nil, err
	}
	return religs, nil
}

func (relig *Religion) Insert() (string, error) {
	relig.ID = bson.NewObjectId()

	if err := db.Religions.Insert(relig); err != nil {
		return "", err
	}
	return relig.ID.Hex(), nil
}

func Find(id string) (*Religion, error) {
	var relig Religion

	if !bson.IsObjectIdHex(id) {
		return &relig, models.ErrInvalidObjectID
	}

	if err := db.Religions.FindId(bson.ObjectIdHex(id)).One(&relig); err != nil {
		return &relig, err
	}
	return &relig, nil
}
