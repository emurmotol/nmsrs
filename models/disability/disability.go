package disability

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type Disability struct {
	ID   int    `schema:"id" json:"id" bson:"id,omitempty"`
	Name string `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Disability, error) {
	disabs := []Disability{}

	if err := db.Disabilities.Find(nil).All(&disabs); err != nil {
		return nil, err
	}
	return disabs, nil
}

func (disab *Disability) Insert() (int, error) {
	if err := db.Disabilities.Insert(disab); err != nil {
		return 0, err
	}
	return disab.ID, nil
}

func FindByID(id int) (*Disability, error) {
	var disab Disability

	if err := db.Disabilities.Find(bson.M{"id": id}).One(&disab); err != nil {
		return &disab, err
	}
	return &disab, nil
}
