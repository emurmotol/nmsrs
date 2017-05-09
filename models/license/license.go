package license

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type License struct {
	ID   int    `schema:"id" json:"id" bson:"id,omitempty"`
	Name string `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]License, error) {
	licns := []License{}

	if err := db.Licenses.Find(nil).Sort("+name").All(&licns); err != nil {
		return nil, err
	}
	return licns, nil
}

func (licn *License) Insert() (int, error) {
	if err := db.Licenses.Insert(licn); err != nil {
		return 0, err
	}
	return licn.ID, nil
}

func FindByID(id int) (*License, error) {
	var licn License

	if err := db.Licenses.Find(bson.M{"id": id}).One(&licn); err != nil {
		return &licn, err
	}
	return &licn, nil
}

func Search(query interface{}) ([]License, error) {
	licns := []License{}

	if err := db.Licenses.Find(query).Sort("+name").All(&licns); err != nil {
		return nil, err
	}
	return licns, nil
}
