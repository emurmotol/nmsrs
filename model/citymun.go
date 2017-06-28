package model

import (
	"github.com/emurmotol/nmsrs/db"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CityMun struct {
	Id       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Code     string        `json:"code" bson:"code"`
	Desc     string        `json:"desc" bson:"desc"`
	PsgcCode string        `json:"psgcCode" bson:"psgcCode"`
	RegCode  string        `json:"regCode" bson:"regCode"`
	ProvCode string        `json:"provCode" bson:"provCode"`
}

type CityMunProv interface{}

func (cityMun *CityMun) Create() *CityMun {
	if err := db.C("cityMuns").Insert(cityMun); err != nil {
		panic(err)
	}
	defer db.Close()
	return cityMun
}

func CityMunById(id bson.ObjectId) *CityMun {
	cityMun := new(CityMun)

	if err := db.C("cityMuns").FindId(id).One(&cityMun); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	return cityMun
}

func (cityMun CityMun) ProvinceIndex(q string) []CityMunProv {
	cityMunProv := []CityMunProv{}
	match := bson.M{
		"$or": []bson.M{
			bson.M{
				"desc": bson.RegEx{
					Pattern: q,
					Options: "i",
				},
			},
			bson.M{
				"province.desc": bson.RegEx{
					Pattern: q,
					Options: "i",
				},
			},
		},
	}

	query := []bson.M{
		bson.M{
			"$lookup": bson.M{
				"from":         "provinces",
				"localField":   "provCode",
				"foreignField": "code",
				"as":           "province",
			},
		},
		bson.M{
			"$match": match,
		},
		bson.M{
			"$sort": bson.M{
				"desc": 1,
			},
		},
	}

	if err := db.C("cityMuns").Pipe(query).All(&cityMunProv); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return cityMunProv
}

func (cityMun *CityMun) BarangayIndex(q string) []Barangay {
	barangays := []Barangay{}
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{
		"$and": []bson.M{
			bson.M{"cityMunCode": cityMun.Code},
			bson.M{"desc": regex},
		},
	}

	if err := db.C("barangays").Find(query).All(&barangays); err != nil {
		if err == mgo.ErrNotFound {
			return nil
		}
		panic(err)
	}
	defer db.Close()
	return barangays
}
