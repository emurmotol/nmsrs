package school

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type School struct {
	ID   int    `schema:"id" json:"id" bson:"id,omitempty"`
	Name string `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]School, error) {
	schs := []School{}

	if err := db.Schools.Find(nil).Sort("+name").All(&schs); err != nil {
		return nil, err
	}
	return schs, nil
}

func (sch *School) Insert() (int, error) {
	if err := db.Schools.Insert(sch); err != nil {
		return 0, err
	}
	return sch.ID, nil
}

func FindByID(id int) (*School, error) {
	var sch School

	if err := db.Schools.Find(bson.M{"id": id}).One(&sch); err != nil {
		return &sch, err
	}
	return &sch, nil
}

func Search(query interface{}) ([]School, error) {
	schs := []School{}

	if err := db.Schools.Find(query).Sort("+name").All(&schs); err != nil {
		return nil, err
	}
	return schs, nil
}
