package sex

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Sex struct {
	ID   bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Sex, error) {
	sexs := []Sex{}

	if err := db.Sexes.Find(bson.M{}).Sort("-name").All(&sexs); err != nil {
		return nil, err
	}
	return sexs, nil
}

func (sex *Sex) Insert() (string, error) {
	sex.ID = bson.NewObjectId()

	if err := db.Sexes.Insert(sex); err != nil {
		return "", err
	}
	return sex.ID.Hex(), nil
}

func Find(id string) (*Sex, error) {
	var sex Sex

	if !bson.IsObjectIdHex(id) {
		return &sex, models.ErrInvalidObjectID
	}

	if err := db.Sexes.FindId(bson.ObjectIdHex(id)).One(&sex); err != nil {
		return &sex, err
	}
	return &sex, nil
}
