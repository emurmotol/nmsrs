package citymunicipality

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type CityMunicipality struct {
	ID           int    `schema:"id" json:"id" bson:"id,omitempty"`
	Code         string `schema:"code" json:"code" bson:"code,omitempty"`
	Desc         string `schema:"desc" json:"desc" bson:"desc,omitempty"`
	PSGCCode     string `schema:"psgc_code" json:"psgc_code" bson:"psgcCode,omitempty"`
	RegionCode   string `schema:"region_code" json:"region_code" bson:"regionCode,omitempty"`
	ProvinceCode string `schema:"province_code" json:"province_code" bson:"provinceCode,omitempty"`
}

func All() ([]CityMunicipality, error) {
	cityMuns := []CityMunicipality{}

	if err := db.CityMunicipalities.Find(nil).Sort("+desc").All(&cityMuns); err != nil {
		return nil, err
	}
	return cityMuns, nil
}

func (cityMun *CityMunicipality) Insert() (int, error) {
	if err := db.CityMunicipalities.Insert(cityMun); err != nil {
		return 0, err
	}
	return cityMun.ID, nil
}

func FindByID(id int) (*CityMunicipality, error) {
	var cityMun CityMunicipality

	if err := db.CityMunicipalities.Find(bson.M{"id": id}).One(&cityMun); err != nil {
		return &cityMun, err
	}
	return &cityMun, nil
}

func Search(query interface{}) ([]interface{}, error) {
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
			"$match": query,
		},
		bson.M{
			"$sort": bson.M{
				"desc": 1,
			},
		},
	}

	if err := db.CityMunicipalities.Pipe(q).All(&cmps); err != nil {
		return nil, err
	}
	return cmps, nil
}
