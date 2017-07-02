package model

import (
	
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Barangay struct {
	Id          bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Code        string        `json:"code" bson:"code"`
	Desc        string        `json:"desc" bson:"desc"`
	RegCode     string        `json:"regCode" bson:"regCode"`
	ProvCode    string        `json:"provCode" bson:"provCode"`
	CityMunCode string        `json:"cityMunCode" bson:"cityMunCode"`
}

func (barangay *Barangay) Create() *Barangay {
	if err := db.C("barangays").Insert(barangay); err != nil {
		panic(err)
	}
	return barangay
}

func BarangayById(id bson.ObjectId) *Barangay {
	barangay := new(Barangay)

	if err := db.C("barangays").FindId(id).One(&barangay); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return barangay
}
