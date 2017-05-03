package country

import (
	"github.com/zneyrl/nmsrs/db"
	"github.com/zneyrl/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Country struct {
	ID   bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Country, error) {
	countries := []Country{}

	if err := db.Countries.Find(bson.M{}).All(&countries); err != nil {
		return nil, err
	}
	return countries, nil
}

func (country *Country) Insert() (string, error) {
	country.ID = bson.NewObjectId()

	if err := db.Countries.Insert(country); err != nil {
		return "", err
	}
	return country.ID.Hex(), nil
}

func Find(id string) (*Country, error) {
	var country Country

	if !bson.IsObjectIdHex(id) {
		return &country, models.ErrInvalidObjectID
	}

	if err := db.Countries.FindId(bson.ObjectIdHex(id)).One(&country); err != nil {
		return &country, err
	}
	return &country, nil
}
