package model

import (
	"github.com/emurmotol/nmsrs/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Language struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Value string        `json:"value" bson:"value"`
}

func (language *Language) Create() *Language {
	if err := db.C("languages").Insert(language); err != nil {
		panic(err)
	}
	defer db.Close()
	return language
}

func (language Language) Index(q string) []Language {
	languages := []Language{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{"name": regex}

	if err := db.C("languages").Find(query).All(&languages); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return languages
}

func LanguageById(id bson.ObjectId) *Language {
	language := new(Language)

	if err := db.C("languages").FindId(id).One(&language); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return language
}
