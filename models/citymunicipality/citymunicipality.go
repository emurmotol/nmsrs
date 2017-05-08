package citymunicipality

import (
	"github.com/emurmotol/nmsrs/db"
	"github.com/emurmotol/nmsrs/models"
	"gopkg.in/mgo.v2/bson"
)

type CityMunicipality struct {
	ID           bson.ObjectId `schema:"id" json:"id" bson:"_id,omitempty"`
	Code         string        `schema:"code" json:"code" bson:"code,omitempty"`
	Desc         string        `schema:"desc" json:"desc" bson:"desc,omitempty"`
	PSGCCode     string        `schema:"psgc_code" json:"psgc_code" bson:"psgcCode,omitempty"`
	RegionCode   string        `schema:"region_code" json:"region_code" bson:"regionCode,omitempty"`
	ProvinceCode string        `schema:"province_code" json:"province_code" bson:"provinceCode,omitempty"`
}

func All() ([]CityMunicipality, error) {
	cityMuns := []CityMunicipality{}

	if err := db.CityMunicipalities.Find(nil).Sort("+desc").All(&cityMuns); err != nil {
		return nil, err
	}
	return cityMuns, nil
}

func (cityMun *CityMunicipality) Insert() (string, error) {
	cityMun.ID = bson.NewObjectId()

	if err := db.CityMunicipalities.Insert(cityMun); err != nil {
		return "", err
	}
	return cityMun.ID.Hex(), nil
}

func FindByID(id string) (*CityMunicipality, error) {
	var cityMun CityMunicipality

	if !bson.IsObjectIdHex(id) {
		return &cityMun, models.ErrInvalidObjectID
	}

	if err := db.CityMunicipalities.FindId(bson.ObjectIdHex(id)).One(&cityMun); err != nil {
		return &cityMun, err
	}
	return &cityMun, nil
}

func FindAllBy(key string, value interface{}) ([]CityMunicipality, error) {
	cityMuns := []CityMunicipality{}

	if err := db.CityMunicipalities.Find(bson.M{key: value}).Sort("+desc").All(&cityMuns); err != nil {
		return cityMuns, err
	}
	return cityMuns, nil
}

func WithProvince(query interface{}) ([]interface{}, error) {
	var cmps []interface{}

	q := []bson.M{
		bson.M{
			"$lookup": bson.M{
				"from":         "provinces",
				"localField":   "provinceCode",
				"foreignField": "code",
				"as":           "province",
			},
		},
		bson.M{
			"$sort": bson.M{
				"desc": 1,
			},
		},
		bson.M{
			"$match": query,
		},
	}

	if err := db.CityMunicipalities.Pipe(q).All(&cmps); err != nil {
		return nil, err
	}
	return cmps, nil
}
