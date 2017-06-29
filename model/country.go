package model

import (
	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Country struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Value string        `json:"value" bson:"value"`
}

func (country *Country) Create() *Country {
	if err := db.C("countries").Insert(country); err != nil {
		panic(err)
	}
	defer db.Close()
	return country
}

func (country Country) Index(q string) []Country {
	countries := []Country{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"name": regex}

	if err := db.C("countries").Find(query).All(&countries); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return countries
}

func CountryById(id bson.ObjectId) *Country {
	country := new(Country)

	if err := db.C("countries").FindId(id).One(&country); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return country
}
