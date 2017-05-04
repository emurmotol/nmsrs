package country

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Country struct {
	ID   bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Country, error) {
	couns := []Country{}

	if err := db.Countries.Find(bson.M{}).Sort("+name").All(&couns); err != nil {
		return nil, err
	}
	return couns, nil
}

func (coun *Country) Insert() (string, error) {
	coun.ID = bson.NewObjectId()

	if err := db.Countries.Insert(coun); err != nil {
		return "", err
	}
	return coun.ID.Hex(), nil
}

func Find(id string) (*Country, error) {
	var coun Country

	if !bson.IsObjectIdHex(id) {
		return &coun, models.ErrInvalidObjectID
	}

	if err := db.Countries.FindId(bson.ObjectIdHex(id)).One(&coun); err != nil {
		return &coun, err
	}
	return &coun, nil
}
