package country

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Country struct {
	ObjectID   bson.ObjectId `schema:"_id" json:"_id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Country, error) {
	couns := []Country{}

	if err := db.Countries.Find(nil).Sort("+name").All(&couns); err != nil {
		return nil, err
	}
	return couns, nil
}

func (coun *Country) Insert() (string, error) {
	coun.ObjectID = bson.NewObjectId()

	if err := db.Countries.Insert(coun); err != nil {
		return "", err
	}
	return coun.ObjectID.Hex(), nil
}

func FindByID(id string) (*Country, error) {
	var coun Country

	if !bson.IsObjectIdHex(id) {
		return &coun, models.ErrInvalidObjectID
	}

	if err := db.Countries.FindId(bson.ObjectIdHex(id)).One(&coun); err != nil {
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
