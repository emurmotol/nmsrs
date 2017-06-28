package model

import (
	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Sex struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}

func (sex *Sex) Create() *Sex {
	if err := db.C("sexes").Insert(sex); err != nil {
		panic(err)
	}
	defer db.Close()
	return sex
}

func Sexes() []Sex {
	sexes := []Sex{}

	if err := db.C("sexes").Find(nil).All(&sexes); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return sexes
}

func SexById(id bson.ObjectId) *Sex {
	sex := new(Sex)

	if err := db.C("sexes").FindId(id).One(&sex); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return sex
}
