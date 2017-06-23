package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Disability struct {
	Id   bson.ObjectId `json:"id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
}

func disabilitySeeder() {
	data := []string{
		"VISUAL IMPAIRMENT",
		"HEARING IMPAIRMENT",
		"SPEECH IMPAIRMENT",
		"PHYSICAL IMPAIRMENT",
		"OTHER",
	}

	for _, name := range data {
		disability := Disability{
			Id:   bson.NewObjectId(),
			Name: strings.ToUpper(name),
		}
		disability.Create()
	}
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

	if err := db.C("disabilities").Find(nil).All(&disabilities); err != mgo.ErrNotFound {
		panic(err)
	} else if err == mgo.ErrNotFound {
		return nil
	}
	defer db.Close()
	return disabilities
}
