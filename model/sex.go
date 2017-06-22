package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/db"

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
	db.C("sexs").Insert(sex)
	defer db.Close()
	return sex
}

func Sexes() []Sex {
	sexes := []Sex{}
	db.C("sexes").Find(nil).All(&sexes)
	defer db.Close()
	return sexes
}
