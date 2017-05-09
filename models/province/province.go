package province

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type Province struct {
	ObjectID         bson.ObjectId `schema:"_id" json:"_id" bson:"_id,omitempty"`
	Code       string        `schema:"code" json:"code" bson:"code,omitempty"`
	Desc       string        `schema:"desc" json:"desc" bson:"desc,omitempty"`
	PSGCCode   string        `schema:"psgc_code" json:"psgc_code" bson:"psgcCode,omitempty"`
	RegionCode string        `schema:"region_code" json:"region_code" bson:"regionCode,omitempty"`
}

func All() ([]Province, error) {
	provs := []Province{}

	if err := db.Provinces.Find(nil).Sort("+desc").All(&provs); err != nil {
		return nil, err
	}
	return provs, nil
}

func (prov *Province) Insert() (string, error) {
	prov.ObjectID = bson.NewObjectId()

	if err := db.Provinces.Insert(prov); err != nil {
		return "", err
	}
	return prov.ObjectID.Hex(), nil
}

func FindByID(id string) (*Province, error) {
	var prov Province

	if !bson.IsObjectIdHex(id) {
		return &prov, models.ErrInvalidObjectID
	}

	if err := db.Provinces.FindId(bson.ObjectIdHex(id)).One(&prov); err != nil {
		return &prov, err
	}
	return &prov, nil
}

func FindAllBy(key string, value interface{}) ([]Province, error) {
	provs := []Province{}

	if err := db.Provinces.Find(bson.M{key: value}).Sort("+desc").All(&provs); err != nil {
		return provs, err
	}
	return provs, nil
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
