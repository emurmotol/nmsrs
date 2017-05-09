package industry

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type Industry struct {
	ID   int    `schema:"id" json:"id" bson:"id,omitempty"`
	Name string `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Industry, error) {
	inds := []Industry{}

	if err := db.Industries.Find(nil).Sort("+name").All(&inds); err != nil {
		return nil, err
	}
	return inds, nil
}

func (ind *Industry) Insert() (int, error) {
	if err := db.Industries.Insert(ind); err != nil {
		return 0, err
	}
	return ind.ID, nil
}

func FindByID(id int) (*Industry, error) {
	var ind Industry

	if err := db.Industries.Find(bson.M{"id": id}).One(&ind); err != nil {
		return &ind, err
	}
	return &ind, nil
}

func Search(query interface{}) ([]Industry, error) {
	inds := []Industry{}

	if err := db.Industries.Find(query).Sort("+name").All(&inds); err != nil {
		return nil, err
	}
	return inds, nil
}
