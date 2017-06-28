package model

import (
	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Religion struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}

func (religion *Religion) Create() *Religion {
	if err := db.C("religions").Insert(religion); err != nil {
		panic(err)
	}
	defer db.Close()
	return religion
}

func (religion Religion) Index(q string) []Religion {
	religions := []Religion{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"name": regex}

	if err := db.C("religions").Find(query).All(&religions); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return religions
}

func ReligionById(id bson.ObjectId) *Religion {
	religion := new(Religion)

	if err := db.C("religions").FindId(id).One(&religion); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return religion
}
