package model

import (
	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Industry struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Value string        `json:"value" bson:"value"`
}

func (industry *Industry) Create() *Industry {
	if err := db.C("industries").Insert(industry); err != nil {
		panic(err)
	}
	defer db.Close()
	return industry
}

func (industry Industry) Index(q string) []Industry {
	industries := []Industry{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"name": regex}

	if err := db.C("industries").Find(query).All(&industries); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return industries
}
