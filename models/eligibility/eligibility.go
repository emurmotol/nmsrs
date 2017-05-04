package eligibility

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Eligibility struct {
	ID   bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Eligibility, error) {
	eligs := []Eligibility{}

	if err := db.Eligibilities.Find(bson.M{}).All(&eligs); err != nil {
		return nil, err
	}
	return eligs, nil
}

func (elig *Eligibility) Insert() (string, error) {
	elig.ID = bson.NewObjectId()

	if err := db.Eligibilities.Insert(elig); err != nil {
		return "", err
	}
	return elig.ID.Hex(), nil
}

func Find(id string) (*Eligibility, error) {
	var elig Eligibility

	if !bson.IsObjectIdHex(id) {
		return &elig, models.ErrInvalidObjectID
	}

	if err := db.Eligibilities.FindId(bson.ObjectIdHex(id)).One(&elig); err != nil {
		return &elig, err
	}
	return &elig, nil
}
