package model

import (
	
	"gopkg.in/mgo.v2/bson"
)

type Region struct {
	Id       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Code     string        `json:"code" bson:"code"`
	Desc     string        `json:"desc" bson:"desc"`
	PsgcCode string        `json:"psgcCode" bson:"psgcCode"`
}

func (region *Region) Create() *Region {
	if err := db.C("regions").Insert(region); err != nil {
		panic(err)
	}
	return region
}
