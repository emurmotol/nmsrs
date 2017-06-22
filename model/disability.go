package model

import (
	"strings"

	"github.com/emurmotol/nmsrs/db"

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
	db.C("disabilities").Insert(disability)
	defer db.Close()
	return disability
}

func Disabilities() []Disability {
	disabilities := []Disability{}
	db.C("disabilities").Find(nil).All(&disabilities)
	defer db.Close()
	return disabilities
}
