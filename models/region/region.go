package region

import (
	"github.com/emurmotol/nmsrs/db"
	"gopkg.in/mgo.v2/bson"
)

type Region struct {
	ID       int    `schema:"id" json:"id" bson:"id,omitempty"`
	Code     string `schema:"code" json:"code" bson:"code,omitempty"`
	Desc     string `schema:"desc" json:"desc" bson:"desc,omitempty"`
	PSGCCode string `schema:"psgc_code" json:"psgc_code" bson:"psgcCode,omitempty"`
}

func All() ([]Region, error) {
	regs := []Region{}

	if err := db.Regions.Find(nil).Sort("+desc").All(&regs); err != nil {
		return nil, err
	}
	return regs, nil
}

func (reg *Region) Insert() (int, error) {
	if err := db.Regions.Insert(reg); err != nil {
		return 0, err
	}
	return reg.ID, nil
}

func FindByID(id int) (*Region, error) {
	var reg Region

	if err := db.Regions.Find(bson.M{"id": id}).One(&reg); err != nil {
		return &reg, err
	}
	return &reg, nil
}

func Search(query interface{}) ([]Region, error) {
	regs := []Region{}

	if err := db.Regions.Find(query).Sort("+desc").All(&regs); err != nil {
		return nil, err
	}
	return regs, nil
}
