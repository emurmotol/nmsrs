package model

import (
	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type License struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Value string        `json:"value" bson:"value"`
}

func (license *License) Create() *License {
	if err := db.C("licenses").Insert(license); err != nil {
		panic(err)
	}
	defer db.Close()
	return license
}

func (license License) Index(q string) []License {
	licenses := []License{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"value": regex}

	if err := db.C("licenses").Find(query).All(&licenses); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return licenses
}

func LicenseById(id bson.ObjectId) *License {
	license := new(License)

	if err := db.C("licenses").FindId(id).One(&license); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return license
}
