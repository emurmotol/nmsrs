package eligibility

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type Eligibility struct {
	ID   int    `schema:"id" json:"id" bson:"id,omitempty"`
	Name string `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Eligibility, error) {
	eligs := []Eligibility{}

	if err := db.Eligibilities.Find(nil).All(&eligs); err != nil {
		return nil, err
	}
	return eligs, nil
}

func (elig *Eligibility) Insert() (int, error) {
	if err := db.Eligibilities.Insert(elig); err != nil {
		return 0, err
	}
	return elig.ID, nil
}

func FindByID(id int) (*Eligibility, error) {
	var elig Eligibility

	if err := db.Eligibilities.Find(bson.M{"id": id}).One(&elig); err != nil {
		return &elig, err
	}
	return &elig, nil
}

func Search(query interface{}) ([]Eligibility, error) {
	eligs := []Eligibility{}

	if err := db.Eligibilities.Find(query).Sort("+name").All(&eligs); err != nil {
		return nil, err
	}
	return eligs, nil
}
