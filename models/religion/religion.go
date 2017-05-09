package religion

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Religion struct {
	ObjectID   bson.ObjectId `schema:"_id" json:"_id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Religion, error) {
	religs := []Religion{}

	if err := db.Religions.Find(nil).Sort("+name").All(&religs); err != nil {
		return nil, err
	}
	return religs, nil
}

func (relig *Religion) Insert() (string, error) {
	relig.ObjectID = bson.NewObjectId()

	if err := db.Religions.Insert(relig); err != nil {
		return "", err
	}
	return relig.ObjectID.Hex(), nil
}

func FindByID(id string) (*Religion, error) {
	var relig Religion

	if !bson.IsObjectIdHex(id) {
		return &relig, models.ErrInvalidObjectID
	}

	if err := db.Religions.FindId(bson.ObjectIdHex(id)).One(&relig); err != nil {
		return &relig, err
	}
	return &relig, nil
}

func Search(query interface{}) ([]Religion, error) {
	religs := []Religion{}

	if err := db.Religions.Find(query).Sort("+name").All(&religs); err != nil {
		return nil, err
	}
	return religs, nil
}
