package country

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type Country struct {
	ID   int    `schema:"id" json:"id" bson:"id,omitempty"`
	Name string `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Country, error) {
	couns := []Country{}

	if err := db.Countries.Find(nil).Sort("+name").All(&couns); err != nil {
		return nil, err
	}
	return couns, nil
}

func (coun *Country) Insert() (int, error) {
	if err := db.Countries.Insert(coun); err != nil {
		return 0, err
	}
	return coun.ID, nil
}

func FindByID(id int) (*Country, error) {
	var coun Country

	if err := db.Countries.Find(bson.M{"id": id}).One(&coun); err != nil {
		return &coun, err
	}
	return &coun, nil
}

func Search(query interface{}) ([]Country, error) {
	couns := []Country{}

	if err := db.Countries.Find(query).Sort("+name").All(&couns); err != nil {
		return nil, err
	}
	return couns, nil
}
