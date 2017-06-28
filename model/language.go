package model

import (
	"encoding/json"
	"io/ioutil"

	"log"

	"github.com/emurmotol/nmsrs/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func languageSeeder() {
	data, err := ioutil.ReadFile("import/languages.json")

	if err != nil {
		panic(err)
	}
	languages := []Language{}

	if err := json.Unmarshal(data, &languages); err != nil {
		panic(err)
	}
	log.Println("languageSeeder: todo")
}

type Language struct {
	Id   bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name string        `json:"name" bson:"name"`
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
