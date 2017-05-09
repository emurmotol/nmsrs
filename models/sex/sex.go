package sex

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type Sex struct {
	ID   int    `schema:"id" json:"id" bson:"id,omitempty"`
	Name string `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Sex, error) {
	sexs := []Sex{}

	if err := db.Sexes.Find(nil).Sort("-name").All(&sexs); err != nil {
		return nil, err
	}
	return sexs, nil
}

func (sex *Sex) Insert() (int, error) {
	if err := db.Sexes.Insert(sex); err != nil {
		return 0, err
	}
	return sex.ID, nil
}

func FindByID(id int) (*Sex, error) {
	var sex Sex

	if err := db.Sexes.Find(bson.M{"id": id}).One(&sex); err != nil {
		return &sex, err
	}
	return &sex, nil
}
