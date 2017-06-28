package model

import (
	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Disability struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
}

func (disability *Disability) Create() *Disability {
	if err := db.C("disabilities").Insert(disability); err != nil {
		panic(err)
	}
	defer db.Close()
	return disability
}

func Disabilities() []Disability {
	disabilities := []Disability{}

	if err := db.C("disabilities").Find(nil).All(&disabilities); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return disabilities
}

func DisabilityById(id bson.ObjectId) *Disability {
	disability := new(Disability)

	if err := db.C("disabilities").FindId(id).One(&disability); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return disability
}
