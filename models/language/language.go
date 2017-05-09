package language

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Language struct {
	ObjectID   bson.ObjectId `schema:"_id" json:"_id" bson:"_id,omitempty"`
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
	lng.ObjectID = bson.NewObjectId()

	if err := db.Languages.Insert(lng); err != nil {
		return "", err
	}
	return lng.ObjectID.Hex(), nil
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

func Search(query interface{}) ([]Language, error) {
	lngs := []Language{}

	if err := db.Languages.Find(query).Sort("+name").All(&lngs); err != nil {
		return nil, err
	}
	return lngs, nil
}
