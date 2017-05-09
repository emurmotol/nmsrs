package province

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type Province struct {
	ID         int    `schema:"id" json:"id" bson:"id,omitempty"`
	Code       string `schema:"code" json:"code" bson:"code,omitempty"`
	Desc       string `schema:"desc" json:"desc" bson:"desc,omitempty"`
	PSGCCode   string `schema:"psgc_code" json:"psgc_code" bson:"psgcCode,omitempty"`
	RegionCode string `schema:"region_code" json:"region_code" bson:"regionCode,omitempty"`
}

func All() ([]Province, error) {
	provs := []Province{}

	if err := db.Provinces.Find(nil).Sort("+desc").All(&provs); err != nil {
		return nil, err
	}
	return provs, nil
}

func (prov *Province) Insert() (int, error) {
	if err := db.Provinces.Insert(prov); err != nil {
		return 0, err
	}
	return prov.ID, nil
}

func FindByID(id int) (*Province, error) {
	var prov Province

	if err := db.Provinces.Find(bson.M{"id": id}).One(&prov); err != nil {
		return &prov, err
	}
	return &prov, nil
}

func FindByCode(code string) *Province {
	var prov Province

	if err := db.Provinces.Find(bson.M{"code": code}).One(&prov); err != nil {
		panic(err)
	}
	return &prov
}

func Search(query interface{}) ([]Province, error) {
	provs := []Province{}

	if err := db.Provinces.Find(query).Sort("+desc").All(&provs); err != nil {
		return nil, err
	}
	return provs, nil
}
