package model

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/emurmotol/nmsrs/db"

	"gopkg.in/mgo.v2/bson"
)

type CityMun struct {
	Id       bson.ObjectId `json:"id" bson:"_id"`
	Code     string        `json:"code" bson:"code"`
	Desc     string        `json:"desc" bson:"desc"`
	PsgcCode string        `json:"psgc_code" bson:"psgcCode"`
	RegCode  string        `json:"reg_code" bson:"regCode"`
	ProvCode string        `json:"prov_code" bson:"provCode"`
}

type CityMunProv interface{}

type RefCityMun struct {
	PsgcCode    string `json:"psgcCode"`
	CityMunDesc string `json:"cityMunDesc"`
	RegCode     string `json:"regCode"`
	ProvCode    string `json:"provCode"`
	CityMunCode string `json:"cityMunCode"`
}

func cityMunSeeder() {
	data, err := ioutil.ReadFile("model/data/refcitymun.json")

	if err != nil {
		panic(err)
	}
	refCityMuns := []RefCityMun{}

	if err := json.Unmarshal(data, &refCityMuns); err != nil {
		panic(err)
	}

	for _, refCityMun := range refCityMuns {
		cityMun := CityMun{
			Id:       bson.NewObjectId(),
			Code:     refCityMun.CityMunCode,
			Desc:     strings.ToUpper(refCityMun.CityMunDesc),
			PsgcCode: refCityMun.PsgcCode,
			RegCode:  refCityMun.RegCode,
			ProvCode: refCityMun.ProvCode,
		}
		cityMun.Create()
	}
}

func (cityMun *CityMun) Create() *CityMun {
	db.C("cityMuns").Insert(cityMun)
	defer db.Close()
	return cityMun
}

func CityMunById(id bson.ObjectId) *CityMun {
	cityMun := new(CityMun)
	db.C("cityMuns").FindId(id).One(&cityMun)
	return cityMun
}

func (cityMun CityMun) ProvinceIndex(q string) []CityMunProv {
	cityMunProv := []CityMunProv{}
	r := make(chan []CityMunProv)
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

	go func() {
		db.C("cityMuns").Pipe(query).All(&cityMunProv)
		defer db.Close()
		r <- cityMunProv
	}()

	cityMunProv = <-r
	close(r)
	return cityMunProv
}

func (cityMun *CityMun) BarangayIndex(q string) []Barangay {
	barangays := []Barangay{}
	r := make(chan []Barangay)
	regex := bson.M{"$regex": bson.RegEx{Pattern: q, Options: "i"}}
	query := bson.M{
		"$and": []bson.M{
			bson.M{"cityMunCode": cityMun.Code},
			bson.M{"desc": regex},
		},
	}

	go func() {
		db.C("barangays").Find(query).All(&barangays)
		defer db.Close()
		r <- barangays
	}()

	barangays = <-r
	close(r)
	return barangays
}
