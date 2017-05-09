package barangay

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Barangay struct {
	ObjectID             bson.ObjectId `schema:"_id" json:"_id" bson:"_id,omitempty"`
	Code                 string        `schema:"code" json:"code" bson:"code,omitempty"`
	Desc                 string        `schema:"desc" json:"desc" bson:"desc,omitempty"`
	RegionCode           string        `schema:"region_code" json:"region_code" bson:"regionCode,omitempty"`
	ProvinceCode         string        `schema:"province_code" json:"province_code" bson:"provinceCode,omitempty"`
	CityMunicipalityCode string        `schema:"city_municipality_code" json:"city_municipality_code" bson:"cityMunicipalityCode,omitempty"`
}

func All() ([]Barangay, error) {
	brgys := []Barangay{}

	if err := db.Barangays.Find(nil).Sort("+desc").All(&brgys); err != nil {
		return nil, err
	}
	return brgys, nil
}

func (brgy *Barangay) Insert() (string, error) {
	brgy.ObjectID = bson.NewObjectId()

	if err := db.Barangays.Insert(brgy); err != nil {
		return "", err
	}
	return brgy.ObjectID.Hex(), nil
}

func FindByID(id string) (*Barangay, error) {
	var brgy Barangay

	if !bson.IsObjectIdHex(id) {
		return &brgy, models.ErrInvalidObjectID
	}

	if err := db.Barangays.FindId(bson.ObjectIdHex(id)).One(&brgy); err != nil {
		return &brgy, err
	}
	return &brgy, nil
}

func Search(query interface{}) ([]Barangay, error) {
	brgys := []Barangay{}

	if err := db.Barangays.Find(query).Sort("+name").All(&brgys); err != nil {
		return nil, err
	}
	return brgys, nil
}
