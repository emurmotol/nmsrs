package course

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type Course struct {
	ID   int    `schema:"id" json:"id" bson:"id,omitempty"`
	Name string `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Course, error) {
	cours := []Course{}

	if err := db.Courses.Find(nil).Sort("+name").All(&cours); err != nil {
		return nil, err
	}
	return cours, nil
}

func (cour *Course) Insert() (int, error) {
	if err := db.Courses.Insert(cour); err != nil {
		return 0, err
	}
	return cour.ID, nil
}

func FindByID(id int) (*Course, error) {
	var cour Course

	if err := db.Courses.Find(bson.M{"id": id}).One(&cour); err != nil {
		return &cour, err
	}
	return &cour, nil
}

func Search(query interface{}) ([]Course, error) {
	cours := []Course{}

	if err := db.Courses.Find(query).Sort("+name").All(&cours); err != nil {
		return nil, err
	}
	return cours, nil
}
