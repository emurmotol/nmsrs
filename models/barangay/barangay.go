package barangay

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type Barangay struct {
	ID                   int    `schema:"id" json:"id" bson:"id,omitempty"`
	Code                 string `schema:"code" json:"code" bson:"code,omitempty"`
	Desc                 string `schema:"desc" json:"desc" bson:"desc,omitempty"`
	RegionCode           string `schema:"region_code" json:"region_code" bson:"regionCode,omitempty"`
	ProvinceCode         string `schema:"province_code" json:"province_code" bson:"provinceCode,omitempty"`
	CityMunicipalityCode string `schema:"city_municipality_code" json:"city_municipality_code" bson:"cityMunicipalityCode,omitempty"`
}

func All() ([]Barangay, error) {
	brgys := []Barangay{}

	if err := db.Barangays.Find(nil).Sort("+desc").All(&brgys); err != nil {
		return nil, err
	}
	return brgys, nil
}

func (brgy *Barangay) Insert() (int, error) {
	if err := db.Barangays.Insert(brgy); err != nil {
		return 0, err
	}
	return brgy.ID, nil
}

func FindByID(id int) (*Barangay, error) {
	var brgy Barangay

	if err := db.Barangays.Find(bson.M{"id": id}).One(&brgy); err != nil {
		return &brgy, err
	}
	return &brgy, nil
}

func Search(query interface{}) ([]Barangay, error) {
	brgys := []Barangay{}

	if err := db.Barangays.Find(query).Sort("+desc").All(&brgys); err != nil {
		return nil, err
	}
	return brgys, nil
}
