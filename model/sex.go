package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Sex struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}

func sexSeeder() {
	data := []string{
		"MALE",
		"FEMALE",
	}

	for _, name := range data {
		sex := Sex{
			Id:   bson.NewObjectId(),
			Name: strings.ToUpper(name),
		}
		sex.Create()
	}
}

func (sex *Sex) Create() *Sex {
	if err := db.C("sexs").Insert(sex); err != nil {
		panic(err)
	}
	defer db.Close()
	return sex
}

func Sexes() []Sex {
	sexes := []Sex{}

	if err := db.C("sexes").Find(nil).All(&sexes); err != mgo.ErrNotFound {
		panic(err)
	} else if err == mgo.ErrNotFound {
		return nil
	}
	defer db.Close()
	return sexes
}
