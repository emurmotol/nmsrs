package religion

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type Religion struct {
	ID   int    `schema:"id" json:"id" bson:"id,omitempty"`
	Name string `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Religion, error) {
	religs := []Religion{}

	if err := db.Religions.Find(nil).Sort("+name").All(&religs); err != nil {
		return nil, err
	}
	return religs, nil
}

func (relig *Religion) Insert() (int, error) {
	if err := db.Religions.Insert(relig); err != nil {
		return 0, err
	}
	return relig.ID, nil
}

func FindByID(id int) (*Religion, error) {
	var relig Religion

	if err := db.Religions.Find(bson.M{"id": id}).One(&relig); err != nil {
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
