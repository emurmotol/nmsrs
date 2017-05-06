package language

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Language struct {
	ID   bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Name string        `schema:"name" json:"name" bson:"name,omitempty"`
}

func All() ([]Language, error) {
	lngs := []Language{}

	if err := db.Languages.Find(nil).Sort("+name").All(&lngs); err != nil {
		return nil, err
	}
	return lngs, nil
}

func (lng *Language) Insert() (string, error) {
	lng.ID = bson.NewObjectId()

	if err := db.Languages.Insert(lng); err != nil {
		return "", err
	}
	return lng.ID.Hex(), nil
}

func FindByID(id string) (*Language, error) {
	var lng Language

	if !bson.IsObjectIdHex(id) {
		return &lng, models.ErrInvalidObjectID
	}

	if err := db.Languages.FindId(bson.ObjectIdHex(id)).One(&lng); err != nil {
		return &lng, err
	}
	return &lng, nil
}
