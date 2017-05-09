package language

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type Language struct {
	ID   int    `schema:"id" json:"id" bson:"id,omitempty"`
	Name string `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Language, error) {
	lngs := []Language{}

	if err := db.Languages.Find(nil).Sort("+name").All(&lngs); err != nil {
		return nil, err
	}
	return lngs, nil
}

func (lng *Language) Insert() (int, error) {
	if err := db.Languages.Insert(lng); err != nil {
		return 0, err
	}
	return lng.ID, nil
}

func FindByID(id int) (*Language, error) {
	var lng Language

	if err := db.Languages.Find(bson.M{"id": id}).One(&lng); err != nil {
		return &lng, err
	}
	return &lng, nil
}

func Search(query interface{}) ([]Language, error) {
	lngs := []Language{}

	if err := db.Languages.Find(query).Sort("+name").All(&lngs); err != nil {
		return nil, err
	}
	return lngs, nil
}
